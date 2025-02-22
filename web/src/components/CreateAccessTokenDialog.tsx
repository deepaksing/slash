import { Button, Input, Modal, ModalDialog, Radio, RadioGroup } from "@mui/joy";
import axios from "axios";
import { useState } from "react";
import { toast } from "react-hot-toast";
import { useTranslation } from "react-i18next";
import useLoading from "../hooks/useLoading";
import useUserStore from "../stores/v1/user";
import Icon from "./Icon";

interface Props {
  onClose: () => void;
  onConfirm?: () => void;
}

const expirationOptions = [
  {
    label: "8 hours",
    value: 3600 * 8,
  },
  {
    label: "1 month",
    value: 3600 * 24 * 30,
  },
  {
    label: "Never",
    value: 0,
  },
];

interface State {
  description: string;
  expiration: number;
}

const CreateAccessTokenDialog: React.FC<Props> = (props: Props) => {
  const { onClose, onConfirm } = props;
  const { t } = useTranslation();
  const currentUser = useUserStore().getCurrentUser();
  const [state, setState] = useState({
    description: "",
    expiration: 3600 * 8,
  });
  const requestState = useLoading(false);

  const setPartialState = (partialState: Partial<State>) => {
    setState({
      ...state,
      ...partialState,
    });
  };

  const handleDescriptionInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPartialState({
      description: e.target.value,
    });
  };

  const handleRoleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPartialState({
      expiration: Number(e.target.value),
    });
  };

  const handleSaveBtnClick = async () => {
    if (!state.description) {
      toast.error("Description is required");
      return;
    }

    try {
      await axios.post(`/api/v2/users/${currentUser.id}/access_tokens`, {
        description: state.description,
        expiresAt: new Date(Date.now() + state.expiration * 1000),
      });

      if (onConfirm) {
        onConfirm();
      }
      onClose();
    } catch (error: any) {
      console.error(error);
      toast.error(error.response.data.message);
    }
  };

  return (
    <Modal open={true}>
      <ModalDialog>
        <div className="flex flex-row justify-between items-center w-80 sm:w-96 mb-4">
          <span className="text-lg font-medium">Create Access Token</span>
          <Button variant="plain" onClick={onClose}>
            <Icon.X className="w-5 h-auto text-gray-600" />
          </Button>
        </div>
        <div>
          <div className="w-full flex flex-col justify-start items-start mb-3">
            <span className="mb-2">
              Description <span className="text-red-600">*</span>
            </span>
            <div className="relative w-full">
              <Input
                className="w-full"
                type="text"
                placeholder="Some description"
                value={state.description}
                onChange={handleDescriptionInputChange}
              />
            </div>
          </div>
          <div className="w-full flex flex-col justify-start items-start mb-3">
            <span className="mb-2">
              Expiration <span className="text-red-600">*</span>
            </span>
            <div className="w-full flex flex-row justify-start items-center text-base">
              <RadioGroup orientation="horizontal" value={state.expiration} onChange={handleRoleInputChange}>
                {expirationOptions.map((option) => (
                  <Radio key={option.value} value={option.value} label={option.label} />
                ))}
              </RadioGroup>
            </div>
          </div>
          <div className="w-full flex flex-row justify-end items-center mt-4 space-x-2">
            <Button color="neutral" variant="plain" disabled={requestState.isLoading} loading={requestState.isLoading} onClick={onClose}>
              {t("common.cancel")}
            </Button>
            <Button color="primary" disabled={requestState.isLoading} loading={requestState.isLoading} onClick={handleSaveBtnClick}>
              {t("common.create")}
            </Button>
          </div>
        </div>
      </ModalDialog>
    </Modal>
  );
};

export default CreateAccessTokenDialog;
