import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useSelector } from 'react-redux';
import { RootState } from '../store';
import SecurityPanel from '../components/settings/SecurityPanel';
import NotificationsPanel from '../components/settings/NotificationsPanel';

type ProfileTab = 'profile' | 'security' | 'notifications';

const Profile: React.FC = () => {
  const navigate = useNavigate();
  const { user } = useSelector((state: RootState) => state.auth);
  const [activeTab, setActiveTab] = useState<ProfileTab>('profile');

  const renderContent = () => {
    switch (activeTab) {
      case 'security':
        return <SecurityPanel onBack={() => setActiveTab('profile')} />;
      case 'notifications':
        return <NotificationsPanel onBack={() => setActiveTab('profile')} />;
      case 'profile':
      default:
        return (
          <div className="flex flex-col gap-8">
            <h2 className="text-2xl font-bold text-[#111618] dark:text-white">Personal Information</h2>
            
            {/* Profile Picture */}
            <div className="flex flex-col items-center gap-4 py-6">
              <div className="bg-center bg-no-repeat aspect-square bg-cover rounded-full size-24" data-alt="User profile picture" style={{backgroundImage: 'url("https://lh3.googleusercontent.com/aida-public/AB6AXuCEnDm4nXPWzfK_8n1ZV5RGr0HFvI-eYrWbYmlNFllajBtfB7wWibe8NguOYZquKumnfgTBmuwBJliBmH8R1J8zELGB6pTFpH6OPzgnXLV4Q9tqp1rfQlK-_wxgAzkWUkRHt9e0xSoz2uUm0wm80crVFjHUkYOFS-ag3UR2nl_8im05tLXROGzVbA-Wc3FbMQ40_2bmKGna2hwUgaVJZ3LonpoO2B3QU-PAYylKIMvN8O7RvmjZqjiBth_xiC44L4kt7L4zIDIo37U")'}}></div>
              <button className="text-primary font-medium hover:underline">
                Change Profile Picture
              </button>
            </div>
            
            {/* Personal Details */}
            <div className="space-y-6">
              <h3 className="text-lg font-semibold text-gray-700 dark:text-gray-300 border-b border-gray-200 dark:border-gray-700 pb-2">Personal Details</h3>
              <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label className="font-medium text-sm text-gray-600 dark:text-gray-300 mb-2 block" htmlFor="firstName">First Name</label>
                  <input 
                    className="form-input w-full rounded-lg border-gray-300 dark:border-gray-600 bg-transparent focus:border-primary focus:ring-primary" 
                    id="firstName" 
                    type="text" 
                    defaultValue={user?.firstName || ''}
                  />
                </div>
                <div>
                  <label className="font-medium text-sm text-gray-600 dark:text-gray-300 mb-2 block" htmlFor="lastName">Last Name</label>
                  <input 
                    className="form-input w-full rounded-lg border-gray-300 dark:border-gray-600 bg-transparent focus:border-primary focus:ring-primary" 
                    id="lastName" 
                    type="text" 
                    defaultValue={user?.lastName || ''}
                  />
                </div>
                <div>
                  <label className="font-medium text-sm text-gray-600 dark:text-gray-300 mb-2 block" htmlFor="email">Email Address</label>
                  <input 
                    className="form-input w-full rounded-lg border-gray-300 dark:border-gray-600 bg-transparent focus:border-primary focus:ring-primary" 
                    id="email" 
                    type="email" 
                    defaultValue={user?.email || ''}
                  />
                </div>
                <div>
                  <label className="font-medium text-sm text-gray-600 dark:text-gray-300 mb-2 block" htmlFor="phone">Phone Number</label>
                  <input 
                    className="form-input w-full rounded-lg border-gray-300 dark:border-gray-600 bg-transparent focus:border-primary focus:ring-primary" 
                    id="phone" 
                    type="tel" 
                    defaultValue={user?.phoneNumber || ''}
                  />
                </div>
              </div>
              <div className="flex justify-end">
                <button className="flex min-w-[84px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-12 px-5 bg-primary text-white gap-2 text-base font-bold leading-normal tracking-[0.015em] hover:bg-primary/90 transition-colors">
                  <span className="truncate">Save Changes</span>
                </button>
              </div>
            </div>
            
            {/* Account Information */}
            <div className="space-y-6">
              <h3 className="text-lg font-semibold text-gray-700 dark:text-gray-300 border-b border-gray-200 dark:border-gray-700 pb-2">Account Information</h3>
              <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label className="font-medium text-sm text-gray-600 dark:text-gray-300 mb-2 block" htmlFor="accountNumber">Account Number</label>
                  <input 
                    className="form-input w-full rounded-lg border-gray-300 dark:border-gray-600 bg-transparent focus:border-primary focus:ring-primary" 
                    id="accountNumber" 
                    type="text" 
                    defaultValue="**** **** **** 1234"
                    readOnly
                  />
                </div>
                <div>
                  <label className="font-medium text-sm text-gray-600 dark:text-gray-300 mb-2 block" htmlFor="accountType">Account Type</label>
                  <input 
                    className="form-input w-full rounded-lg border-gray-300 dark:border-gray-600 bg-transparent focus:border-primary focus:ring-primary" 
                    id="accountType" 
                    type="text" 
                    defaultValue="Personal"
                    readOnly
                  />
                </div>
                <div>
                  <label className="font-medium text-sm text-gray-600 dark:text-gray-300 mb-2 block" htmlFor="memberSince">Member Since</label>
                  <input 
                    className="form-input w-full rounded-lg border-gray-300 dark:border-gray-600 bg-transparent focus:border-primary focus:ring-primary" 
                    id="memberSince" 
                    type="text" 
                    defaultValue="January 15, 2023"
                    readOnly
                  />
                </div>
                <div>
                  <label className="font-medium text-sm text-gray-600 dark:text-gray-300 mb-2 block" htmlFor="lastLogin">Last Login</label>
                  <input 
                    className="form-input w-full rounded-lg border-gray-300 dark:border-gray-600 bg-transparent focus:border-primary focus:ring-primary" 
                    id="lastLogin" 
                    type="text" 
                    defaultValue="October 28, 2025 at 3:45 PM"
                    readOnly
                  />
                </div>
              </div>
            </div>
          </div>
        );
    }
  };

  return (
    <div className="relative flex h-auto min-h-screen w-full flex-col group/design-root overflow-x-hidden">
      <main className="w-full max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
        <div className="flex flex-col gap-4 mb-8">
          <div className="flex items-center gap-4">
            <button 
              onClick={() => navigate('/dashboard')}
              className="flex h-10 w-10 items-center justify-center rounded-full text-gray-500 hover:bg-gray-100 dark:hover:bg-gray-700 transition-colors"
            >
              <span className="material-symbols-outlined text-2xl">arrow_back</span>
            </button>
            <h1 className="text-4xl font-black leading-tight tracking-[-0.033em] text-[#111618] dark:text-white">Profile</h1>
          </div>
          <p className="text-base font-normal leading-normal text-gray-500 dark:text-gray-400">Manage your personal information and account settings.</p>
        </div>
        
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
          {/* Left Column: Navigation */}
          <div className="lg:col-span-1 flex flex-col gap-6">
            <div className="bg-white dark:bg-background-dark rounded-xl border border-gray-200 dark:border-gray-700 shadow-sm p-6">
              <div className="flex flex-col gap-4">
                <button 
                  className={`flex items-center gap-3 w-full p-3 rounded-lg text-left transition-colors ${
                    activeTab === 'profile' 
                      ? 'bg-primary/10 dark:bg-primary/20 text-primary' 
                      : 'hover:bg-gray-100 dark:hover:bg-gray-800'
                  }`}
                  onClick={() => setActiveTab('profile')}
                >
                  <span className="material-symbols-outlined text-xl">person</span>
                  <span className="font-medium">Profile</span>
                </button>
                <button 
                  className={`flex items-center gap-3 w-full p-3 rounded-lg text-left transition-colors ${
                    activeTab === 'security' 
                      ? 'bg-primary/10 dark:bg-primary/20 text-primary' 
                      : 'hover:bg-gray-100 dark:hover:bg-gray-800'
                  }`}
                  onClick={() => setActiveTab('security')}
                >
                  <span className="material-symbols-outlined text-xl">security</span>
                  <span className="font-medium">Security</span>
                </button>
                <button 
                  className={`flex items-center gap-3 w-full p-3 rounded-lg text-left transition-colors ${
                    activeTab === 'notifications' 
                      ? 'bg-primary/10 dark:bg-primary/20 text-primary' 
                      : 'hover:bg-gray-100 dark:hover:bg-gray-800'
                  }`}
                  onClick={() => setActiveTab('notifications')}
                >
                  <span className="material-symbols-outlined text-xl">notifications</span>
                  <span className="font-medium">Notifications</span>
                </button>
              </div>
            </div>
            
            <div className="bg-white dark:bg-background-dark rounded-xl border border-gray-200 dark:border-gray-700 shadow-sm p-6">
              <div className="flex flex-col gap-4">
                <button 
                  className="flex items-center gap-3 w-full p-3 rounded-lg text-left text-red-500 hover:bg-red-50 dark:hover:bg-red-900/20 transition-colors"
                  onClick={() => {
                    localStorage.removeItem('token');
                    navigate('/login');
                  }}
                >
                  <span className="material-symbols-outlined text-xl">logout</span>
                  <span className="font-medium">Log Out</span>
                </button>
              </div>
            </div>
          </div>
          
          {/* Right Column: Content */}
          <div className="lg:col-span-2 bg-white dark:bg-background-dark rounded-xl border border-gray-200 dark:border-gray-700 shadow-sm p-8">
            {renderContent()}
          </div>
        </div>
      </main>
    </div>
  );
};

export default Profile;