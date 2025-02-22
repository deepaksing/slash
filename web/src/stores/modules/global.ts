import { createSlice, PayloadAction } from "@reduxjs/toolkit";

type State = {
  workspaceProfile: WorkspaceProfile;
};

const globalSlice = createSlice({
  name: "global",
  initialState: {} as State,
  reducers: {
    setGlobalState: (_, action: PayloadAction<State>) => {
      return action.payload;
    },
  },
});

export const { setGlobalState } = globalSlice.actions;

export default globalSlice.reducer;
