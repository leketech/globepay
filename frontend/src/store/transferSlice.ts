import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { transferService, TransferRequest, Transfer } from '../services/transfer.service';

interface TransferState {
  transfers: Transfer[];
  currentTransfer: Transfer | null;
  loading: boolean;
  error: string | null;
}

const initialState: TransferState = {
  transfers: [],
  currentTransfer: null,
  loading: false,
  error: null,
};

export const createTransfer = createAsyncThunk(
  'transfer/create',
  async (data: TransferRequest, { rejectWithValue }) => {
    try {
      const response = await transferService.createTransfer(data);
      return response;
    } catch (error: any) {
      return rejectWithValue(error.response?.data?.message || 'Failed to create transfer');
    }
  }
);

export const getTransfers = createAsyncThunk(
  'transfer/getAll',
  async (_, { rejectWithValue }) => {
    try {
      const response = await transferService.getTransfers();
      return response.transfers;
    } catch (error: any) {
      return rejectWithValue(error.response?.data?.message || 'Failed to fetch transfers');
    }
  }
);

export const getTransfer = createAsyncThunk(
  'transfer/get',
  async (id: string, { rejectWithValue }) => {
    try {
      const response = await transferService.getTransfer(id);
      return response;
    } catch (error: any) {
      return rejectWithValue(error.response?.data?.message || 'Failed to fetch transfer');
    }
  }
);

const transferSlice = createSlice({
  name: 'transfer',
  initialState,
  reducers: {
    clearError: (state) => {
      state.error = null;
    },
    clearCurrentTransfer: (state) => {
      state.currentTransfer = null;
    },
  },
  extraReducers: (builder) => {
    builder
      // Create transfer
      .addCase(createTransfer.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(createTransfer.fulfilled, (state, action: PayloadAction<Transfer>) => {
        state.loading = false;
        state.transfers.unshift(action.payload);
        state.currentTransfer = action.payload;
      })
      .addCase(createTransfer.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload as string;
      })
      // Get transfers
      .addCase(getTransfers.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(getTransfers.fulfilled, (state, action: PayloadAction<Transfer[]>) => {
        state.loading = false;
        state.transfers = action.payload;
      })
      .addCase(getTransfers.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload as string;
      })
      // Get transfer
      .addCase(getTransfer.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(getTransfer.fulfilled, (state, action: PayloadAction<Transfer>) => {
        state.loading = false;
        state.currentTransfer = action.payload;
      })
      .addCase(getTransfer.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload as string;
      });
  },
});

export const { clearError, clearCurrentTransfer } = transferSlice.actions;
export default transferSlice.reducer;