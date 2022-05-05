import { Action, configureStore, ThunkAction } from '@reduxjs/toolkit';
import travelPlanReducer from '../features/travelPlan/travelPlanSlice';

export const store = configureStore({
  reducer: {
    travelPlan: travelPlanReducer,
  },
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;
