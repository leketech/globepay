export interface User {
  id: string;
  email: string;
  firstName: string;
  lastName: string;
  phoneNumber?: string;
  dateOfBirth?: string;
  country?: string;
  kycStatus: string;
  accountStatus: string;
  createdAt: string;
  updatedAt: string;
}

export interface Account {
  id: string;
  userId: string;
  currency: string;
  balance: number;
  accountNumber: string;
  accountType: string;
  status: string;
  createdAt: string;
  updatedAt: string;
}

export interface Transaction {
  id: string;
  userId: string;
  type: 'credit' | 'debit';
  status: string;
  amount: number;
  currency: string;
  sourceAccount?: string;
  destAccount?: string;
  fee?: number;
  exchangeRate?: number;
  description: string;
  reference: string;
  processedAt?: string;
  failureReason?: string;
  createdAt: string;
  updatedAt: string;
}

export interface Transfer {
  id: string;
  userId: string;
  recipientName: string;
  recipientEmail: string;
  recipientCountry: string;
  recipientBankName: string;
  recipientAccountNo: string;
  recipientSwiftCode: string;
  sourceCurrency: string;
  destCurrency: string;
  sourceAmount: number;
  destAmount: number;
  fee: number;
  exchangeRate: number;
  purpose: string;
  status: string;
  transactionId: string;
  estimatedArrival: string;
  createdAt: string;
  updatedAt: string;
}

export interface AuthState {
  isAuthenticated: boolean;
  user: User | null;
  token: string | null;
  accounts: Account[];
  loading: boolean;
  error: string | null;
}

export interface TransferState {
  transfers: Transfer[];
  currentTransfer: Transfer | null;
  loading: boolean;
  error: string | null;
}

export interface TransactionState {
  transactions: Transaction[];
  currentTransaction: Transaction | null;
  loading: boolean;
  error: string | null;
}

export interface RootState {
  auth: AuthState;
  transfer: TransferState;
  transaction: TransactionState;
}
