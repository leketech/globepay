import React, { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { RootState } from '../../store';
import { createTransfer } from '../../store/transferSlice';
import { TransferRequest } from '../../services/transfer.service';

export const TransferForm: React.FC = () => {
  const dispatch = useDispatch();
  const { accounts } = useSelector((state: RootState) => state.auth);
  const { loading, error } = useSelector((state: RootState) => state.transfer);

  const [formData, setFormData] = useState({
    fromAccount: '',
    toAccount: '',
    amount: '',
    currency: 'USD',
    description: '',
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    
    // Create the transfer request object with the correct interface
    const transferData: TransferRequest = {
      recipientName: 'Recipient', // This should be updated with actual recipient name
      recipientCountry: 'US', // This should be updated with actual recipient country
      recipientBankName: 'Bank', // This should be updated with actual bank name
      recipientAccountNo: formData.toAccount,
      recipientSwiftCode: 'SWIFT', // This should be updated with actual SWIFT code
      sourceCurrency: formData.currency,
      destCurrency: formData.currency, // This should be updated with destination currency
      sourceAmount: parseFloat(formData.amount),
      purpose: formData.description,
    };

    dispatch(createTransfer(transferData) as any);
  };

  return (
    <div className="bg-white shadow rounded-lg p-6">
      <h2 className="text-lg font-medium text-gray-900 mb-4">Send Money</h2>
      <form onSubmit={handleSubmit} className="space-y-6">
        <div>
          <label htmlFor="fromAccount" className="block text-sm font-medium text-gray-700">
            From Account
          </label>
          <select
            id="fromAccount"
            name="fromAccount"
            value={formData.fromAccount}
            onChange={handleChange}
            className="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm rounded-md"
            required
          >
            <option value="">Select account</option>
            {accounts.map((account) => (
              <option key={account.id} value={account.id}>
                {account.accountNumber} ({account.currency}) - {account.balance.toLocaleString('en-US', {
                  style: 'currency',
                  currency: account.currency,
                })}
              </option>
            ))}
          </select>
        </div>

        <div>
          <label htmlFor="toAccount" className="block text-sm font-medium text-gray-700">
            To Account
          </label>
          <input
            type="text"
            id="toAccount"
            name="toAccount"
            value={formData.toAccount}
            onChange={handleChange}
            placeholder="Enter recipient account number"
            className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
            required
          />
        </div>

        <div>
          <label htmlFor="amount" className="block text-sm font-medium text-gray-700">
            Amount
          </label>
          <div className="mt-1 relative rounded-md shadow-sm">
            <input
              type="number"
              id="amount"
              name="amount"
              value={formData.amount}
              onChange={handleChange}
              placeholder="0.00"
              step="0.01"
              min="0"
              className="block w-full pr-12 border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
              required
            />
            <div className="absolute inset-y-0 right-0 flex items-center">
              <label htmlFor="currency" className="sr-only">
                Currency
              </label>
              <select
                id="currency"
                name="currency"
                value={formData.currency}
                onChange={handleChange}
                className="h-full py-0 pl-2 pr-7 border-transparent bg-transparent text-gray-500 focus:ring-blue-500 focus:border-blue-500 sm:text-sm rounded-md"
              >
                <option>USD</option>
                <option>EUR</option>
                <option>GBP</option>
              </select>
            </div>
          </div>
        </div>

        <div>
          <label htmlFor="description" className="block text-sm font-medium text-gray-700">
            Description
          </label>
          <input
            type="text"
            id="description"
            name="description"
            value={formData.description}
            onChange={handleChange}
            placeholder="Enter transfer description"
            className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
          />
        </div>

        {error && (
          <div className="rounded-md bg-red-50 p-4">
            <div className="flex">
              <div className="flex-shrink-0">
                <svg className="h-5 w-5 text-red-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                  <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clipRule="evenodd" />
                </svg>
              </div>
              <div className="ml-3">
                <h3 className="text-sm font-medium text-red-800">{error}</h3>
              </div>
            </div>
          </div>
        )}

        <div>
          <button
            type="submit"
            disabled={loading}
            className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50"
          >
            {loading ? 'Processing...' : 'Send Money'}
          </button>
        </div>
      </form>
    </div>
  );
};