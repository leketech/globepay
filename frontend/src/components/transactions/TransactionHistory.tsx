import React from 'react';
import { useSelector } from 'react-redux';
import { RootState } from '../../store';

export const TransactionHistory: React.FC = () => {
  const { transactions } = useSelector((state: RootState) => state.transaction);

  const formatCurrency = (amount: number, currency: string) => {
    return new Intl.NumberFormat('en-US', {
      style: 'currency',
      currency,
    }).format(amount);
  };

  const formatDate = (dateString: string) => {
    const date = new Date(dateString);
    return date.toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
    });
  };

  return (
    <div className="bg-white shadow rounded-lg overflow-hidden">
      <div className="px-6 py-4 border-b border-gray-200">
        <h3 className="text-lg font-medium text-gray-900">Recent Transactions</h3>
      </div>
      <ul className="divide-y divide-gray-200">
        {transactions.map((transaction) => (
          <li key={transaction.id} className="px-6 py-4">
            <div className="flex items-center justify-between">
              <div className="flex items-center">
                <div className="flex-shrink-0">
                  <div
                    className={`h-10 w-10 rounded-full flex items-center justify-center ${
                      transaction.type === 'credit' ? 'bg-green-100' : 'bg-red-100'
                    }`}
                  >
                    <span
                      className={`text-sm font-medium ${
                        transaction.type === 'credit' ? 'text-green-800' : 'text-red-800'
                      }`}
                    >
                      {transaction.type === 'credit' ? '+' : '-'}
                    </span>
                  </div>
                </div>
                <div className="ml-4">
                  <div className="text-sm font-medium text-gray-900">{transaction.description}</div>
                  <div className="text-sm text-gray-500">{formatDate(transaction.createdAt)}</div>
                </div>
              </div>
              <div
                className={`text-sm font-medium ${
                  transaction.type === 'credit' ? 'text-green-600' : 'text-red-600'
                }`}
              >
                {transaction.type === 'credit' ? '+' : '-'}
                {formatCurrency(transaction.amount, transaction.currency)}
              </div>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
};
