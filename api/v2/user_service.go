package v2

import (
	"context"

	"github.com/boojack/slash/api/auth"
	apiv2pb "github.com/boojack/slash/proto/gen/api/v2"
	storepb "github.com/boojack/slash/proto/gen/store"
	"github.com/boojack/slash/store"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserService struct {
	apiv2pb.UnimplementedUserServiceServer

	Secret string
	Store  *store.Store
}

// NewUserService creates a new UserService.
func NewUserService(secret string, store *store.Store) *UserService {
	return &UserService{
		Secret: secret,
		Store:  store,
	}
}

func (s *UserService) ListUsers(ctx context.Context, _ *apiv2pb.ListUsersRequest) (*apiv2pb.ListUsersResponse, error) {
	users, err := s.Store.ListUsers(ctx, &store.FindUser{})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list users: %v", err)
	}

	userMessages := []*apiv2pb.User{}
	for _, user := range users {
		userMessages = append(userMessages, convertUserFromStore(user))
	}
	response := &apiv2pb.ListUsersResponse{
		Users: userMessages,
	}
	return response, nil
}

func (s *UserService) GetUser(ctx context.Context, request *apiv2pb.GetUserRequest) (*apiv2pb.GetUserResponse, error) {
	user, err := s.Store.GetUser(ctx, &store.FindUser{
		ID: &request.Id,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find user: %v", err)
	}
	if user == nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	userMessage := convertUserFromStore(user)
	response := &apiv2pb.GetUserResponse{
		User: userMessage,
	}
	return response, nil
}

func (s *UserService) CreateUser(ctx context.Context, request *apiv2pb.CreateUserRequest) (*apiv2pb.CreateUserResponse, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %v", err)
	}

	user, err := s.Store.CreateUser(ctx, &store.User{
		Email:        request.User.Email,
		Nickname:     request.User.Nickname,
		Role:         store.RoleUser,
		PasswordHash: string(passwordHash),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}
	response := &apiv2pb.CreateUserResponse{
		User: convertUserFromStore(user),
	}
	return response, nil
}

func (s *UserService) DeleteUser(ctx context.Context, request *apiv2pb.DeleteUserRequest) (*apiv2pb.DeleteUserResponse, error) {
	userID := ctx.Value(UserIDContextKey).(int32)
	if userID == request.Id {
		return nil, status.Errorf(codes.InvalidArgument, "cannot delete self")
	}

	err := s.Store.DeleteUser(ctx, &store.DeleteUser{
		ID: request.Id,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete user: %v", err)
	}
	response := &apiv2pb.DeleteUserResponse{}
	return response, nil
}

func (s *UserService) ListUserAccessTokens(ctx context.Context, request *apiv2pb.ListUserAccessTokensRequest) (*apiv2pb.ListUserAccessTokensResponse, error) {
	userID := ctx.Value(UserIDContextKey).(int32)
	if userID != request.Id {
		return nil, status.Errorf(codes.PermissionDenied, "Permission denied")
	}

	userAccessTokens, err := s.Store.GetUserAccessTokens(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list access tokens: %v", err)
	}

	accessTokens := []*apiv2pb.UserAccessToken{}
	for _, userAccessToken := range userAccessTokens {
		claims := &auth.ClaimsMessage{}
		_, err := jwt.ParseWithClaims(userAccessToken.AccessToken, claims, func(t *jwt.Token) (any, error) {
			if t.Method.Alg() != jwt.SigningMethodHS256.Name {
				return nil, errors.Errorf("unexpected access token signing method=%v, expect %v", t.Header["alg"], jwt.SigningMethodHS256)
			}
			if kid, ok := t.Header["kid"].(string); ok {
				if kid == "v1" {
					return []byte(s.Secret), nil
				}
			}
			return nil, errors.Errorf("unexpected access token kid=%v", t.Header["kid"])
		})
		if err != nil {
			// If the access token is invalid or expired, just ignore it.
			continue
		}

		userAccessToken := &apiv2pb.UserAccessToken{
			AccessToken: userAccessToken.AccessToken,
			Description: userAccessToken.Description,
			IssuedAt:    timestamppb.New(claims.IssuedAt.Time),
		}
		if claims.ExpiresAt != nil {
			userAccessToken.ExpiresAt = timestamppb.New(claims.ExpiresAt.Time)
		}
		accessTokens = append(accessTokens, userAccessToken)
	}

	// Sort by issued time in descending order.
	slices.SortFunc(accessTokens, func(i, j *apiv2pb.UserAccessToken) bool {
		return i.IssuedAt.Seconds > j.IssuedAt.Seconds
	})
	response := &apiv2pb.ListUserAccessTokensResponse{
		AccessTokens: accessTokens,
	}
	return response, nil
}

func (s *UserService) CreateUserAccessToken(ctx context.Context, request *apiv2pb.CreateUserAccessTokenRequest) (*apiv2pb.CreateUserAccessTokenResponse, error) {
	userID := ctx.Value(UserIDContextKey).(int32)
	if userID != request.Id {
		return nil, status.Errorf(codes.PermissionDenied, "Permission denied")
	}

	user, err := s.Store.GetUser(ctx, &store.FindUser{
		ID: &userID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}
	if user == nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	accessToken, err := auth.GenerateAccessToken(user.Email, user.ID, request.UserAccessToken.ExpiresAt.AsTime(), s.Secret)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate access token: %v", err)
	}

	claims := &auth.ClaimsMessage{}
	_, err = jwt.ParseWithClaims(accessToken, claims, func(t *jwt.Token) (any, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Name {
			return nil, errors.Errorf("unexpected access token signing method=%v, expect %v", t.Header["alg"], jwt.SigningMethodHS256)
		}
		if kid, ok := t.Header["kid"].(string); ok {
			if kid == "v1" {
				return []byte(s.Secret), nil
			}
		}
		return nil, errors.Errorf("unexpected access token kid=%v", t.Header["kid"])
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to parse access token: %v", err)
	}

	// Upsert the access token to user setting store.
	if err := s.UpsertAccessTokenToStore(ctx, user, accessToken, request.UserAccessToken.Description); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to upsert access token to store: %v", err)
	}

	userAccessToken := &apiv2pb.UserAccessToken{
		AccessToken: accessToken,
		Description: request.UserAccessToken.Description,
		IssuedAt:    timestamppb.New(claims.IssuedAt.Time),
	}
	if claims.ExpiresAt != nil {
		userAccessToken.ExpiresAt = timestamppb.New(claims.ExpiresAt.Time)
	}
	response := &apiv2pb.CreateUserAccessTokenResponse{
		AccessToken: userAccessToken,
	}
	return response, nil
}

func (s *UserService) DeleteUserAccessToken(ctx context.Context, request *apiv2pb.DeleteUserAccessTokenRequest) (*apiv2pb.DeleteUserAccessTokenResponse, error) {
	userID := ctx.Value(UserIDContextKey).(int32)
	if userID != request.Id {
		return nil, status.Errorf(codes.PermissionDenied, "Permission denied")
	}

	userAccessTokens, err := s.Store.GetUserAccessTokens(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list access tokens: %v", err)
	}
	updatedUserAccessTokens := []*storepb.AccessTokensUserSetting_AccessToken{}
	for _, userAccessToken := range userAccessTokens {
		if userAccessToken.AccessToken == request.AccessToken {
			continue
		}
		updatedUserAccessTokens = append(updatedUserAccessTokens, userAccessToken)
	}
	if _, err := s.Store.UpsertUserSetting(ctx, &storepb.UserSetting{
		UserId: userID,
		Key:    storepb.UserSettingKey_USER_SETTING_ACCESS_TOKENS,
		Value: &storepb.UserSetting_AccessTokensUserSetting{
			AccessTokensUserSetting: &storepb.AccessTokensUserSetting{
				AccessTokens: updatedUserAccessTokens,
			},
		},
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to upsert user setting: %v", err)
	}

	return &apiv2pb.DeleteUserAccessTokenResponse{}, nil
}

func (s *UserService) UpsertAccessTokenToStore(ctx context.Context, user *store.User, accessToken, description string) error {
	userAccessTokens, err := s.Store.GetUserAccessTokens(ctx, user.ID)
	if err != nil {
		return errors.Wrap(err, "failed to get user access tokens")
	}
	userAccessToken := storepb.AccessTokensUserSetting_AccessToken{
		AccessToken: accessToken,
		Description: description,
	}
	userAccessTokens = append(userAccessTokens, &userAccessToken)
	if _, err := s.Store.UpsertUserSetting(ctx, &storepb.UserSetting{
		UserId: user.ID,
		Key:    storepb.UserSettingKey_USER_SETTING_ACCESS_TOKENS,
		Value: &storepb.UserSetting_AccessTokensUserSetting{
			AccessTokensUserSetting: &storepb.AccessTokensUserSetting{
				AccessTokens: userAccessTokens,
			},
		},
	}); err != nil {
		return errors.Wrap(err, "failed to upsert user setting")
	}
	return nil
}

func convertUserFromStore(user *store.User) *apiv2pb.User {
	return &apiv2pb.User{
		Id:        int32(user.ID),
		RowStatus: convertRowStatusFromStore(user.RowStatus),
		CreatedTs: user.CreatedTs,
		UpdatedTs: user.UpdatedTs,
		Role:      convertUserRoleFromStore(user.Role),
		Email:     user.Email,
		Nickname:  user.Nickname,
	}
}

func convertUserRoleFromStore(role store.Role) apiv2pb.Role {
	switch role {
	case store.RoleAdmin:
		return apiv2pb.Role_ADMIN
	case store.RoleUser:
		return apiv2pb.Role_USER
	default:
		return apiv2pb.Role_ROLE_UNSPECIFIED
	}
}
