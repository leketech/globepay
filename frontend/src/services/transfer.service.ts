import { transferApi } from './api';
import { Transfer as TransferType } from '../types';

export interface TransferRequest {
  recipientName: string;
  recipientCountry: string;
  recipientBankName: string;
  recipientAccountNo: string;
  recipientSwiftCode: string;
  sourceCurrency: string;
  destCurrency: string;
  sourceAmount: number;
  purpose: string;
}

export interface Transfer {
  id: string;
  userId: string;
  recipientName: string;
  recipientCountry: string;
  sourceAmount: number;
  destAmount: number;
  sourceCurrency: string;
  destCurrency: string;
  exchangeRate: number;
  fee: number;
  status: string;
  estimatedArrival: string;
  createdAt: string;
}

export interface ExchangeRateResponse {
  fromCurrency: string;
  toCurrency: string;
  rate: number;
  fee: number;
  amount: number;
  convertedAmount: number;
  timestamp: string;
}

export const transferService = {
  async getExchangeRate(from: string, to: string, amount: number): Promise<ExchangeRateResponse> {
    try {
      console.log('Getting exchange rate from backend API...', { from, to, amount });

      // Validate inputs
      if (!from || !to || amount <= 0) {
        throw new Error('Invalid input parameters for exchange rate calculation');
      }

      // Call the backend API endpoint for exchange rates
      const response = await fetch(`/api/v1/exchange-rates?from=${from}&to=${to}&amount=${amount}`);

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data: ExchangeRateResponse = await response.json();
      console.log('Backend API response:', data);

      // Validate response
      if (!data || !isFinite(data.rate) || !isFinite(data.convertedAmount)) {
        throw new Error('Invalid response from backend API');
      }

      return data;
    } catch (error) {
      console.error('Failed to get exchange rate from backend API:', error);
      // Fallback to a mock implementation with reasonable defaults
      const mockRate =
        from === 'USD' && to === 'EUR'
          ? 0.85
          : from === 'EUR' && to === 'USD'
          ? 1.18
          : from === 'USD' && to === 'GBP'
          ? 0.75
          : from === 'GBP' && to === 'USD'
          ? 1.33
          : from === 'USD' && to === 'NGN'
          ? 1580.0
          : 1.0; // Default 1:1 rate

      const fee = this.calculateFee(amount);
      const convertedAmount = (amount - fee) * mockRate;

      return {
        fromCurrency: from,
        toCurrency: to,
        rate: mockRate,
        fee: fee,
        amount: amount,
        convertedAmount: convertedAmount,
        timestamp: new Date().toISOString(),
      };
    }
  },

  calculateFee(amount: number): number {
    let fee = 0;
    if (amount <= 100) {
      fee = 2.99; // Fixed fee for small amounts
    } else if (amount <= 1000) {
      fee = amount * 0.01; // 1% fee for medium amounts
    } else {
      fee = amount * 0.005; // 0.5% fee for large amounts
    }

    // Validate fee
    if (!isFinite(fee)) {
      fee = 0;
    }

    return fee;
  },

  async createTransfer(data: TransferRequest): Promise<TransferType> {
    // Get token from localStorage
    const token = localStorage.getItem('token') || '';
    
    // Map TransferRequest to Transfer for the API call
    const transferData: any = {
      recipientName: data.recipientName,
      recipientCountry: data.recipientCountry,
      recipientBankName: data.recipientBankName,
      recipientAccountNo: data.recipientAccountNo,
      recipientSwiftCode: data.recipientSwiftCode,
      sourceCurrency: data.sourceCurrency,
      destCurrency: data.destCurrency,
      sourceAmount: data.sourceAmount,
      purpose: data.purpose,
      // These fields will be populated by the backend
      id: '',
      userId: '',
      recipientEmail: '',
      destAmount: 0,
      fee: 0,
      exchangeRate: 0,
      status: 'pending',
      transactionId: '',
      estimatedArrival: '',
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString(),
    };
    
    // Using the transferApi.createTransfer function directly
    return transferApi.createTransfer(token, transferData);
  },

  async getTransfer(id: string): Promise<Transfer> {
    // Get token from localStorage
    const token = localStorage.getItem('token') || '';
    // Using the transferApi.getTransfer function directly
    return transferApi.getTransfer(token, id);
  },

  async getTransfers(page = 1, limit = 10): Promise<{ transfers: Transfer[]; total: number }> {
    // Get token from localStorage
    const token = localStorage.getItem('token') || '';
    // Using the transferApi.getTransfers function directly
    try {
      const result = await transferApi.getTransfers(token, page, limit);
      console.log('Transfers API response:', result);
      return result;
    } catch (error) {
      console.error('Error fetching transfers:', error);
      // Return empty data structure on error
      return { transfers: [], total: 0 };
    }
  },

  async cancelTransfer(id: string): Promise<void> {
    // Get token from localStorage
    const token = localStorage.getItem('token') || '';
    // Using the transferApi.cancelTransfer function directly
    return transferApi.cancelTransfer(token, id);
  },
};
