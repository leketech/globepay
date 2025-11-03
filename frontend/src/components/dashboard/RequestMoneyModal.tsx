import React, { useState } from 'react';

interface RequestMoneyModalProps {
  isOpen: boolean;
  onClose: () => void;
  onRequestMoney: (amount: number, recipient: string, isLink: boolean) => void;
}

const RequestMoneyModal: React.FC<RequestMoneyModalProps> = ({ isOpen, onClose, onRequestMoney }) => {
  const [amount, setAmount] = useState('');
  const [recipient, setRecipient] = useState('');
  const [requestType, setRequestType] = useState<'user' | 'link'>('user');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    const amountValue = parseFloat(amount);
    if (isNaN(amountValue) || amountValue <= 0) {
      alert('Please enter a valid amount');
      return;
    }
    if (requestType === 'user' && !recipient) {
      alert('Please enter a recipient');
      return;
    }
    onRequestMoney(amountValue, recipient, requestType === 'link');
    onClose();
  };

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
      <div className="bg-white dark:bg-gray-800 rounded-xl shadow-lg w-full max-w-md">
        <div className="p-6">
          <div className="flex justify-between items-center mb-4">
            <h2 className="text-xl font-bold text-gray-900 dark:text-white">Request Money</h2>
            <button 
              onClick={onClose}
              className="text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200"
            >
              <span className="material-symbols-outlined">close</span>
            </button>
          </div>

          <form onSubmit={handleSubmit}>
            <div className="mb-4">
              <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                Amount
              </label>
              <div className="relative">
                <span className="absolute inset-y-0 left-0 pl-3 flex items-center text-gray-500 dark:text-gray-400">
                  $
                </span>
                <input
                  type="number"
                  value={amount}
                  onChange={(e) => setAmount(e.target.value)}
                  className="pl-8 w-full rounded-lg border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-white py-2 px-3 focus:outline-none focus:ring-2 focus:ring-primary/50"
                  placeholder="0.00"
                  step="0.01"
                  min="0"
                  required
                />
              </div>
            </div>

            <div className="mb-6">
              <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                Request Type
              </label>
              <div className="grid grid-cols-2 gap-3">
                <button
                  type="button"
                  onClick={() => setRequestType('user')}
                  className={`p-4 rounded-lg border ${
                    requestType === 'user'
                      ? 'border-primary bg-primary/10 dark:bg-primary/20'
                      : 'border-gray-300 dark:border-gray-600 hover:bg-gray-50 dark:hover:bg-gray-700'
                  }`}
                >
                  <div className="flex flex-col items-center">
                    <span className="material-symbols-outlined text-2xl mb-1">person</span>
                    <span className="text-sm font-medium">From User</span>
                  </div>
                </button>
                <button
                  type="button"
                  onClick={() => setRequestType('link')}
                  className={`p-4 rounded-lg border ${
                    requestType === 'link'
                      ? 'border-primary bg-primary/10 dark:bg-primary/20'
                      : 'border-gray-300 dark:border-gray-600 hover:bg-gray-50 dark:hover:bg-gray-700'
                  }`}
                >
                  <div className="flex flex-col items-center">
                    <span className="material-symbols-outlined text-2xl mb-1">link</span>
                    <span className="text-sm font-medium">Payment Link</span>
                  </div>
                </button>
              </div>
            </div>

            {requestType === 'user' ? (
              <div className="mb-4">
                <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                  Recipient Name or ID
                </label>
                <input
                  type="text"
                  value={recipient}
                  onChange={(e) => setRecipient(e.target.value)}
                  className="w-full rounded-lg border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-white py-2 px-3 focus:outline-none focus:ring-2 focus:ring-primary/50"
                  placeholder="Enter recipient name or ID"
                  required
                />
              </div>
            ) : (
              <div className="mb-6 p-4 bg-blue-50 dark:bg-blue-900/30 rounded-lg">
                <p className="text-sm text-blue-800 dark:text-blue-200">
                  After submitting, a payment link will be generated that you can share with anyone.
                  They can use this link to send you money.
                </p>
              </div>
            )}

            <div className="flex gap-3 mt-6">
              <button
                type="button"
                onClick={onClose}
                className="flex-1 py-2 px-4 rounded-lg border border-gray-300 dark:border-gray-600 text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors"
              >
                Cancel
              </button>
              <button
                type="submit"
                className="flex-1 py-2 px-4 rounded-lg bg-primary text-white hover:bg-primary/90 transition-colors"
              >
                {requestType === 'user' ? 'Request Money' : 'Generate Link'}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
};

export default RequestMoneyModal;