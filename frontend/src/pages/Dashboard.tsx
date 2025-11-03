import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { useSelector } from 'react-redux';
import { RootState } from '../store';
import { transferService } from '../services/transfer.service';
import { Transfer } from '../services/transfer.service';
import AddMoneyModal from '../components/dashboard/AddMoneyModal';
import RequestMoneyModal from '../components/dashboard/RequestMoneyModal';

const Dashboard: React.FC = () => {
  const navigate = useNavigate();
  const { user } = useSelector((state: RootState) => state.auth);
  const [balance, setBalance] = useState(1250.55);
  const [recentTransfers, setRecentTransfers] = useState<Transfer[]>([]);
  const [loading, setLoading] = useState(true);
  const [searchTerm, setSearchTerm] = useState('');
  const [isAddMoneyModalOpen, setIsAddMoneyModalOpen] = useState(false);
  const [isRequestMoneyModalOpen, setIsRequestMoneyModalOpen] = useState(false);

  useEffect(() => {
    loadDashboardData();
  }, []);

  const loadDashboardData = async () => {
    try {
      const { transfers } = await transferService.getTransfers(1, 5);
      setRecentTransfers(transfers);
    } catch (error) {
      console.error('Failed to load dashboard data:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleAddMoney = (amount: number, method: 'card' | 'bank') => {
    // In a real app, this would call an API to add money to the user's account
    setBalance((prevBalance) => prevBalance + amount);
    console.log(`Added $${amount} via ${method}`);
    // Show success message
    alert(`Successfully added $${amount} to your account!`);
  };

  const handleRequestMoney = (amount: number, recipient: string, isLink: boolean) => {
    if (isLink) {
      // Generate payment link
      const link = `${window.location.origin}/pay?amount=${amount}&recipient=${encodeURIComponent(
        user?.firstName || 'User'
      )}`;
      navigator.clipboard.writeText(link);
      alert(`Payment link generated and copied to clipboard!\n\n${link}`);
    } else {
      // Send request to specific user
      console.log(`Requested $${amount} from ${recipient}`);
      alert(`Money request for $${amount} sent to ${recipient}!`);
    }
  };

  if (loading) {
    return (
      <div className="min-h-screen bg-background-light dark:bg-background-dark flex items-center justify-center">
        <div className="text-lg text-text-primary-light dark:text-text-primary-dark">
          Loading...
        </div>
      </div>
    );
  }

  // Filter transfers based on search term
  const filteredTransfers = recentTransfers.filter(
    (transfer) =>
      transfer.recipientName.toLowerCase().includes(searchTerm.toLowerCase()) ||
      transfer.id.toLowerCase().includes(searchTerm.toLowerCase())
  );

  return (
    <div className="relative flex h-auto min-h-screen w-full flex-col group/design-root overflow-x-hidden">
      {/* Add Money Modal */}
      <AddMoneyModal
        isOpen={isAddMoneyModalOpen}
        onClose={() => setIsAddMoneyModalOpen(false)}
        onAddMoney={handleAddMoney}
      />

      {/* Request Money Modal */}
      <RequestMoneyModal
        isOpen={isRequestMoneyModalOpen}
        onClose={() => setIsRequestMoneyModalOpen(false)}
        onRequestMoney={handleRequestMoney}
      />

      <div className="layout-container flex h-full grow flex-col">
        <main className="flex-1 w-full max-w-7xl mx-auto p-4 sm:p-6 lg:p-10">
          <div className="flex flex-col lg:flex-row gap-8">
            {/* Main Content (Left Column) */}
            <div className="w-full lg:w-2/3 flex flex-col gap-8">
              <p className="text-4xl font-black leading-tight tracking-[-0.033em]">
                Welcome back, {user?.firstName || 'User'}!
              </p>

              {/* Balance Card */}
              <div className="p-6 @container bg-content-light dark:bg-content-dark rounded-xl shadow-sm">
                <div className="flex flex-col items-stretch justify-start gap-4 sm:flex-row sm:items-center">
                  <div className="flex w-full min-w-72 grow flex-col items-stretch justify-center gap-1">
                    <p className="text-text-secondary-light dark:text-text-secondary-dark text-sm font-normal leading-normal">
                      Available Balance
                    </p>
                    <p className="text-3xl font-bold leading-tight tracking-[-0.015em]">
                      ${balance.toFixed(2)} USD
                    </p>
                  </div>
                  <div className="flex items-center justify-start sm:justify-end gap-3">
                    <button
                      onClick={() => setIsAddMoneyModalOpen(true)}
                      className="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-10 px-6 bg-primary text-white text-sm font-medium leading-normal hover:bg-primary/90 transition-colors"
                    >
                      <span className="truncate">Add Money</span>
                    </button>
                  </div>
                </div>
              </div>

              {/* Recent Activity Section */}
              <div className="flex flex-col gap-4">
                <div className="flex flex-col sm:flex-row justify-between items-baseline gap-4 px-1">
                  <h2 className="text-[22px] font-bold leading-tight tracking-[-0.015em]">
                    Recent Activity
                  </h2>
                  <button
                    onClick={() => navigate('/history')}
                    className="text-sm font-medium text-primary hover:underline"
                  >
                    View All
                  </button>
                </div>

                <div className="px-0 py-3">
                  <label className="flex flex-col min-w-40 h-12 w-full">
                    <div className="flex w-full flex-1 items-stretch rounded-lg h-full">
                      <div className="flex items-center justify-center pl-4 rounded-l-lg bg-gray-100 dark:bg-gray-800">
                        <span className="material-symbols-outlined text-text-secondary-light dark:text-text-secondary-dark">
                          search
                        </span>
                      </div>
                      <input
                        className="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-r-lg text-text-primary-light dark:text-text-primary-dark focus:outline-0 focus:ring-2 focus:ring-primary/50 border-none bg-gray-100 dark:bg-gray-800 h-full placeholder:text-text-secondary-light dark:placeholder:text-text-secondary-dark px-4 pl-2 text-base font-normal leading-normal"
                        placeholder="Search transactions..."
                        value={searchTerm}
                        onChange={(e) => setSearchTerm(e.target.value)}
                      />
                    </div>
                  </label>
                </div>

                {/* Transaction List */}
                <div className="flex flex-col bg-content-light dark:bg-content-dark rounded-xl shadow-sm overflow-hidden">
                  <ul className="divide-y divide-gray-200/80 dark:divide-gray-700/80">
                    {filteredTransfers.length > 0 ? (
                      filteredTransfers.map((transfer) => (
                        <li
                          key={transfer.id}
                          className="p-4 flex items-center justify-between gap-4 hover:bg-background-light dark:hover:bg-background-dark/50 transition-colors"
                        >
                          <div className="flex items-center gap-4">
                            <div
                              className="bg-center bg-no-repeat aspect-square bg-cover rounded-full size-10"
                              style={{
                                backgroundImage:
                                  'url("https://lh3.googleusercontent.com/aida-public/AB6AXuAuHX-31OOI9vwcX77AokzoeEia8aJjisXNFT_pCzqi0hwVsCiLdeT6eNUnqYPhvx7d16l0SX-V793iW9Wa31bQICo1qLIBAQ-SDrNsaKoMLAtvgvI2poJzFvb3lGujYwgOJX5USjXBsavayMzjWdt-lAxC8KCFr7r7dCo1H5Eg9d50piz6TsjkpmMaJGtuOz6XV_CwiAqK1B68fFPvhzXM4QpIxAdzeU-CtJDzqhm0FO5dQX_vSYLpQ2W329UBsmsmWG0-Pr1d8lY")',
                              }}
                            ></div>
                            <div>
                              <p className="font-semibold">Sent to {transfer.recipientName}</p>
                              <p className="text-sm text-text-secondary-light dark:text-text-secondary-dark">
                                {new Date(transfer.createdAt).toLocaleDateString()}
                              </p>
                            </div>
                          </div>
                          <p className="font-semibold">
                            {transfer.status === 'completed' ? '-' : ''} $
                            {transfer.sourceAmount.toFixed(2)}
                          </p>
                        </li>
                      ))
                    ) : (
                      <li className="p-4 text-center text-text-secondary-light dark:text-text-secondary-dark">
                        No transactions found
                      </li>
                    )}
                  </ul>
                </div>
              </div>
            </div>

            {/* Sidebar (Right Column) */}
            <aside className="w-full lg:w-1/3 flex flex-col gap-6">
              <div className="bg-content-light dark:bg-content-dark rounded-xl shadow-sm p-6">
                <h3 className="text-lg font-bold mb-4">Quick Actions</h3>
                <div className="flex flex-col gap-4">
                  <button
                    onClick={() => navigate('/transfer')}
                    className="w-full flex items-center justify-center gap-2 h-12 px-6 bg-primary text-white rounded-lg font-medium hover:bg-primary/90 transition-colors"
                  >
                    <span className="material-symbols-outlined">arrow_upward</span>
                    <span>Send Money</span>
                  </button>
                  <button
                    onClick={() => setIsRequestMoneyModalOpen(true)}
                    className="w-full flex items-center justify-center gap-2 h-12 px-6 bg-gray-100 dark:bg-gray-800 text-text-primary-light dark:text-text-primary-dark rounded-lg font-medium hover:bg-gray-200 dark:hover:bg-gray-700 transition-colors"
                  >
                    <span className="material-symbols-outlined">arrow_downward</span>
                    <span>Request Money</span>
                  </button>
                </div>
              </div>

              <div className="bg-content-light dark:bg-content-dark rounded-xl shadow-sm p-6">
                <h3 className="text-lg font-bold mb-4">Recent Payees</h3>
                <div className="flex flex-col gap-4">
                  <div className="flex items-center justify-between">
                    <div className="flex items-center gap-3">
                      <div
                        className="bg-center bg-no-repeat aspect-square bg-cover rounded-full size-10"
                        style={{
                          backgroundImage:
                            'url("https://lh3.googleusercontent.com/aida-public/AB6AXuB_1zEUTEf596tqyxfQkzuZjVa2gBP36XsrLDAr3IxKXT79t-S6HI6a_N-PZ5C4FNDX4M0XxlpYMdAMU7U4eKDebdngasVYtJtbJhktAjq7h4rj2DlOfnOpKTa89a-uGcm_VNqFToOXM25FpcKN_CWk4poFVo9rA3FeAa982WIBWzKfzdfwVJP03lFnBIJYlSPFl9q0bO_SVgfW6wyessAF0XF1gMn4PtVEVLHH-3jVQz4LOQRVwadQmZPJBsgGTagE2W-ZjWMTeF0")',
                        }}
                      ></div>
                      <span className="font-medium">Jane Doe</span>
                    </div>
                    <button className="text-primary text-sm font-medium hover:underline">
                      Pay
                    </button>
                  </div>
                  <div className="flex items-center justify-between">
                    <div className="flex items-center gap-3">
                      <div
                        className="bg-center bg-no-repeat aspect-square bg-cover rounded-full size-10"
                        style={{
                          backgroundImage:
                            'url("https://lh3.googleusercontent.com/aida-public/AB6AXuCu3BzvK2mGWdstU9m5XlCRzn_f7i1iqBEI-tHo4IXJJHQ3fwKRv8TkFnj-5-rr-_DbqK3q3oxBIxBIwMer3rdIgCL0ZTONBkN9UF7dDjW0lWOv7WEDFoGUkJC2HX2d_E1FKy9fscVPUXWV2brTcbx5LMXaEaydB5hfj-cLSzew2YTkWQRHSoVVVhZi6wWcQDR0jqGlc_wUr8Vmc_8W-Ku_FVE5CyTZ-g1ggFeqa8vnElV64HkHOxrhFUHqeUi03ByxWUCi3-UIBM8")',
                        }}
                      ></div>
                      <span className="font-medium">Mike Ross</span>
                    </div>
                    <button className="text-primary text-sm font-medium hover:underline">
                      Pay
                    </button>
                  </div>
                  <div className="flex items-center justify-between">
                    <div className="flex items-center gap-3">
                      <div className="bg-blue-100 dark:bg-blue-900 text-blue-600 dark:text-blue-300 rounded-full size-10 flex items-center justify-center font-bold">
                        AC
                      </div>
                      <span className="font-medium">Acme Corp</span>
                    </div>
                    <button className="text-primary text-sm font-medium hover:underline">
                      Pay
                    </button>
                  </div>
                </div>
              </div>
            </aside>
          </div>
        </main>
      </div>
    </div>
  );
};

export default Dashboard;
