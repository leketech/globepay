import React from 'react';
import { useSelector } from 'react-redux';
import { RootState } from '../../store';
import { Account } from '../../types';

export const AccountSummary: React.FC = () => {
  const { accounts } = useSelector((state: RootState) => state.auth);

  return (
    <div className="bg-white shadow rounded-lg p-6">
      <h2 className="text-lg font-medium text-gray-900 mb-4">Account Summary</h2>
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        {accounts.map((account: Account) => (
          <div key={account.id} className="border rounded-lg p-4">
            <div className="flex justify-between items-center">
              <div>
                <p className="text-sm font-medium text-gray-500">{account.currency}</p>
                <p className="text-2xl font-semibold text-gray-900">
                  {account.balance.toLocaleString('en-US', {
                    style: 'currency',
                    currency: account.currency,
                  })}
                </p>
              </div>
              <div className="bg-blue-100 rounded-full p-2">
                <span className="text-blue-800 font-medium">{account.accountNumber}</span>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};
