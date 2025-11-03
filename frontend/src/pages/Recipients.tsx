import React, { useState } from 'react';
import RecipientForm from '../components/recipient/RecipientForm';

interface Recipient {
  id: number;
  name: string;
  country: string;
  currency: string;
  payoutMethod: string;
  accountDetails: string;
  bankName?: string;
  accountNumber?: string;
  sortCode?: string;
  iban?: string;
  swiftCode?: string;
}

const Recipients: React.FC = () => {
  const [searchTerm, setSearchTerm] = useState('');
  const [selectedRecipient, setSelectedRecipient] = useState<number | null>(1);
  const [showAddForm, setShowAddForm] = useState(false);
  const [recipients, setRecipients] = useState<Recipient[]>([
    {
      id: 1,
      name: 'Jane Doe',
      country: 'United Kingdom',
      currency: 'GBP',
      payoutMethod: 'Bank',
      accountDetails: '•••• 1234',
      bankName: 'Barclays Bank',
      accountNumber: '12345678',
      sortCode: '12-34-56'
    },
    {
      id: 2,
      name: 'John Smith',
      country: 'Germany',
      currency: 'EUR',
      payoutMethod: 'Mobile Wallet',
      accountDetails: '•••• 5678'
    },
    {
      id: 3,
      name: 'Emily Jones',
      country: 'Canada',
      currency: 'CAD',
      payoutMethod: 'Bank',
      accountDetails: '•••• 5678'
    }
  ]);

  const filteredRecipients = recipients.filter(recipient => 
    recipient.name.toLowerCase().includes(searchTerm.toLowerCase())
  );

  const handleEditRecipient = (id: number) => {
    setSelectedRecipient(id);
  };

  const handleDeleteRecipient = (id: number) => {
    setRecipients(recipients.filter(recipient => recipient.id !== id));
    if (selectedRecipient === id) {
      setSelectedRecipient(recipients.length > 1 ? recipients[0].id : null);
    }
  };

  const handleAddRecipient = () => {
    setShowAddForm(true);
    setSelectedRecipient(null);
  };

  const handleSaveNewRecipient = (recipientData: any) => {
    const newId = Math.max(...recipients.map(r => r.id), 0) + 1;
    const recipientToAdd = {
      id: newId,
      name: recipientData.name,
      country: recipientData.country,
      currency: recipientData.currency,
      payoutMethod: 'Bank',
      accountDetails: `•••• ${recipientData.accountNumber.slice(-4)}`,
      bankName: recipientData.bankName,
      accountNumber: recipientData.accountNumber,
      sortCode: recipientData.sortCode,
      iban: recipientData.iban,
      swiftCode: recipientData.swiftCode
    };
    
    setRecipients([...recipients, recipientToAdd]);
    setSelectedRecipient(newId);
    setShowAddForm(false);
  };

  const handleUpdateRecipient = (recipientData: any) => {
    setRecipients(recipients.map(recipient => 
      recipient.id === selectedRecipient 
        ? {
            ...recipient,
            name: recipientData.name,
            country: recipientData.country,
            currency: recipientData.currency,
            accountDetails: `•••• ${recipientData.accountNumber.slice(-4)}`,
            bankName: recipientData.bankName,
            accountNumber: recipientData.accountNumber,
            sortCode: recipientData.sortCode,
            iban: recipientData.iban,
            swiftCode: recipientData.swiftCode
          }
        : recipient
    ));
    setShowAddForm(false);
  };

  const handleCancelAdd = () => {
    setShowAddForm(false);
    setSelectedRecipient(recipients.length > 0 ? recipients[0].id : null);
  };

  return (
    <div className="relative flex h-auto min-h-screen w-full flex-col group/design-root overflow-x-hidden">
      <main className="w-full max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
        <div className="flex flex-col gap-4 mb-8">
          <h1 className="text-4xl font-black leading-tight tracking-[-0.033em] text-[#111618] dark:text-white">Manage Recipients</h1>
          <p className="text-base font-normal leading-normal text-gray-500 dark:text-gray-400">Add, edit, or delete your saved recipients.</p>
        </div>
        
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
          {/* Left Column: Recipient List */}
          <div className="lg:col-span-1 flex flex-col gap-6">
            <div className="relative">
              <span className="material-symbols-outlined absolute left-4 top-1/2 -translate-y-1/2 text-gray-400 dark:text-gray-500">search</span>
              <input 
                className="form-input w-full rounded-lg border-0 bg-gray-100 dark:bg-background-dark h-12 placeholder:text-gray-500 dark:placeholder:text-gray-400 pl-12 pr-4 text-base font-normal text-[#111618] dark:text-gray-200 focus:ring-2 focus:ring-primary focus:ring-offset-2 focus:ring-offset-background-light dark:focus:ring-offset-background-dark" 
                placeholder="Search by name..." 
                type="text"
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
              />
            </div>
            
            <div className="space-y-3">
              {filteredRecipients.map(recipient => (
                <div 
                  key={recipient.id}
                  className={`flex items-center gap-4 p-3 rounded-lg border transition-all cursor-pointer ${
                    selectedRecipient === recipient.id 
                      ? 'bg-primary/10 dark:bg-primary/20 border-primary' 
                      : 'bg-white dark:bg-background-dark border-gray-200 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-800'
                  }`}
                  onClick={() => {
                    setSelectedRecipient(recipient.id);
                    setShowAddForm(false);
                  }}
                >
                  <div className="flex-shrink-0 w-10 h-10 rounded-full bg-primary/10 dark:bg-primary/20 flex items-center justify-center">
                    <span className="text-primary font-medium">
                      {recipient.name.charAt(0)}
                    </span>
                  </div>
                  <div className="flex-grow">
                    <p className="text-base font-semibold leading-normal line-clamp-1 text-[#111618] dark:text-white">{recipient.name}</p>
                    <p className="text-sm font-normal leading-normal line-clamp-2 text-gray-500 dark:text-gray-400">
                      {recipient.bankName ? `${recipient.bankName} •••• ${recipient.accountNumber?.slice(-4)}` : recipient.accountDetails}
                    </p>
                  </div>
                  <div className="flex items-center gap-1">
                    <button 
                      className="p-2 rounded-lg text-gray-500 dark:text-gray-400 hover:bg-gray-100 dark:hover:bg-gray-700 hover:text-primary dark:hover:text-primary transition-colors"
                      onClick={(e) => {
                        e.stopPropagation();
                        handleEditRecipient(recipient.id);
                        setShowAddForm(false);
                      }}
                    >
                      <span className="material-symbols-outlined text-xl">edit</span>
                    </button>
                    <button 
                      className="p-2 rounded-lg text-gray-500 dark:text-gray-400 hover:bg-destructive/10 hover:text-destructive transition-colors"
                      onClick={(e) => {
                        e.stopPropagation();
                        handleDeleteRecipient(recipient.id);
                      }}
                    >
                      <span className="material-symbols-outlined text-xl">delete</span>
                    </button>
                  </div>
                </div>
              ))}
            </div>
            
            <button 
              className="flex min-w-[84px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-12 px-5 w-full bg-primary/20 text-primary gap-2 text-base font-bold leading-normal tracking-[0.015em] hover:bg-primary/30 transition-colors"
              onClick={handleAddRecipient}
            >
              <span className="material-symbols-outlined">add</span>
              <span className="truncate">Add New Recipient</span>
            </button>
          </div>
          
          {/* Right Column: Form Panel */}
          <div className="lg:col-span-2 bg-white dark:bg-background-dark rounded-xl border border-gray-200 dark:border-gray-700 shadow-sm p-8">
            {showAddForm ? (
              <RecipientForm 
                onSubmit={handleSaveNewRecipient}
                onCancel={handleCancelAdd}
              />
            ) : selectedRecipient ? (
              <RecipientForm 
                onSubmit={handleUpdateRecipient}
                onCancel={() => setShowAddForm(false)}
                initialData={{
                  name: recipients.find(r => r.id === selectedRecipient)?.name || '',
                  bankName: recipients.find(r => r.id === selectedRecipient)?.bankName || '',
                  accountNumber: recipients.find(r => r.id === selectedRecipient)?.accountNumber || '',
                  sortCode: recipients.find(r => r.id === selectedRecipient)?.sortCode || '',
                  iban: recipients.find(r => r.id === selectedRecipient)?.iban || '',
                  swiftCode: recipients.find(r => r.id === selectedRecipient)?.swiftCode || '',
                  country: recipients.find(r => r.id === selectedRecipient)?.country || 'United Kingdom',
                  currency: recipients.find(r => r.id === selectedRecipient)?.currency || 'GBP'
                }}
              />
            ) : (
              <div className="flex flex-col items-center justify-center h-full py-12">
                <span className="material-symbols-outlined text-6xl text-gray-300 dark:text-gray-600 mb-4">group</span>
                <h3 className="text-xl font-semibold text-gray-700 dark:text-gray-300 mb-2">No Recipient Selected</h3>
                <p className="text-gray-500 dark:text-gray-400 mb-6">Select a recipient from the list or add a new one</p>
                <button 
                  className="flex min-w-[84px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-12 px-5 bg-primary text-white gap-2 text-base font-bold leading-normal tracking-[0.015em] hover:bg-primary/90 transition-colors"
                  onClick={handleAddRecipient}
                >
                  <span className="material-symbols-outlined">add</span>
                  <span className="truncate">Add New Recipient</span>
                </button>
              </div>
            )}
          </div>
        </div>
      </main>
    </div>
  );
};

export default Recipients;