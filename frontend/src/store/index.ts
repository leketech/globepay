import { configureStore } from '@reduxjs/toolkit';
import authReducer from './authSlice';
import transferReducer from './transferSlice';
import transactionReducer from './transactionSlice';

export const store = configureStore({
  reducer: {
    auth: authReducer,
    transfer: transferReducer,
    transaction: transactionReducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: false,
    }),
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;