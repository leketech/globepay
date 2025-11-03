import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { transferService, Transfer } from '../services/transfer.service';

const History: React.FC = () => {
  const navigate = useNavigate();
  const [transfers, setTransfers] = useState<Transfer[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    loadTransfers();
  }, []);

  const loadTransfers = async () => {
    try {
      const { transfers } = await transferService.getTransfers();
      setTransfers(transfers);
    } catch (err) {
      setError('Failed to load transfers');
      console.error('Error loading transfers:', err);
    } finally {
      setLoading(false);
    }
  };

  if (loading) {
    return (
      <div className="min-h-screen bg-background-light dark:bg-background-dark flex items-center justify-center">
        <div className="text-lg text-gray-600 dark:text-gray-400">Loading...</div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-background-light dark:bg-background-dark p-4">
      <div className="max-w-md mx-auto">
        <div className="flex items-center justify-between mb-6">
          <button onClick={() => navigate('/dashboard')} className="text-primary">
            <span className="material-symbols-outlined">arrow_back</span>
          </button>
          <h1 className="text-xl font-bold text-center text-slate-800 dark:text-white">
            Transaction History
          </h1>
          <div></div>
        </div>

        {error && (
          <div className="mb-4 bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded relative">
            {error}
          </div>
        )}

        <div className="space-y-4">
          {transfers.length === 0 ? (
            <div className="text-center py-12">
              <p className="text-gray-500 dark:text-gray-400">No transactions found</p>
            </div>
          ) : (
            transfers.map((transfer) => (
              <div key={transfer.id} className="bg-white dark:bg-[#192b33] rounded-lg p-4 shadow">
                <div className="flex justify-between items-center mb-2">
                  <div className="flex items-center">
                    <div className="bg-gray-200 dark:bg-[#233c48] rounded-full p-2 mr-3">
                      <span className="material-symbols-outlined text-slate-800 dark:text-white">
                        send
                      </span>
                    </div>
                    <div>
                      <p className="font-medium text-slate-800 dark:text-white">
                        {transfer.recipientName}
                      </p>
                      <p className="text-sm text-gray-500 dark:text-gray-400">
                        {new Date(transfer.createdAt).toLocaleDateString()}
                      </p>
                    </div>
                  </div>
                  <div className="text-right">
                    <p
                      className={`font-bold ${
                        transfer.status === 'completed'
                          ? 'text-green-500'
                          : 'text-slate-800 dark:text-white'
                      }`}
                    >
                      {transfer.sourceAmount ? `-${transfer.sourceAmount.toFixed(2)}` : 'N/A'}{' '}
                      {transfer.sourceCurrency || 'USD'}
                    </p>
                    <p className="text-sm text-gray-500 dark:text-gray-400">
                      {transfer.status || 'pending'}
                    </p>
                  </div>
                </div>

                <div className="flex justify-between text-sm mt-3">
                  <span className="text-gray-500 dark:text-gray-400">
                    Fee: {transfer.fee ? transfer.fee.toFixed(2) : '0.00'}{' '}
                    {transfer.sourceCurrency || 'USD'}
                  </span>
                  <span className="text-gray-500 dark:text-gray-400">
                    Rate: {transfer.exchangeRate ? transfer.exchangeRate.toFixed(4) : 'N/A'}
                  </span>
                </div>
              </div>
            ))
          )}
        </div>
      </div>
    </div>
  );
};

export default History;
