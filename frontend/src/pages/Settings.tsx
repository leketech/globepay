import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../hooks/useAuth';
import { useDispatch } from 'react-redux';
import { logout } from '../store/authSlice';
import type { AppDispatch } from '../store';
import SecurityPanel from '../components/settings/SecurityPanel';
import NotificationsPanel from '../components/settings/NotificationsPanel';

type SettingsTab = 'appearance' | 'security' | 'notifications';

const Settings: React.FC = () => {
  const navigate = useNavigate();
  const dispatch = useDispatch<AppDispatch>();
  useAuth();
  const [activeTab, setActiveTab] = useState<SettingsTab>('appearance');
  const [theme, setTheme] = useState<'light' | 'dark' | 'system'>('system');
  const [emailNotifications, setEmailNotifications] = useState(true);
  const [pushNotifications, setPushNotifications] = useState(false);
  const [smsNotifications, setSmsNotifications] = useState(false);

  useEffect(() => {
    // Check saved preferences
    const savedTheme = localStorage.getItem('theme') as 'light' | 'dark' | 'system' | null;
    const savedEmailNotifications = localStorage.getItem('emailNotifications');
    const savedPushNotifications = localStorage.getItem('pushNotifications');
    const savedSmsNotifications = localStorage.getItem('smsNotifications');
    
    if (savedTheme) {
      setTheme(savedTheme);
      applyTheme(savedTheme);
    } else {
      // Default to system preference
      setTheme('system');
      applyTheme('system');
    }
    
    if (savedEmailNotifications) {
      setEmailNotifications(savedEmailNotifications === 'true');
    }
    
    if (savedPushNotifications) {
      setPushNotifications(savedPushNotifications === 'true');
    }
    
    if (savedSmsNotifications) {
      setSmsNotifications(savedSmsNotifications === 'true');
    }
    
    // Listen for system theme changes
    const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
    const handleSystemThemeChange = (e: MediaQueryListEvent) => {
      const currentTheme = localStorage.getItem('theme') as 'light' | 'dark' | 'system' | null;
      // Only apply system theme changes if user has selected 'system'
      if (!currentTheme || currentTheme === 'system') {
        if (e.matches) {
          document.documentElement.classList.add('dark');
        } else {
          document.documentElement.classList.remove('dark');
        }
      }
    };
    
    mediaQuery.addEventListener('change', handleSystemThemeChange);
    
    return () => {
      mediaQuery.removeEventListener('change', handleSystemThemeChange);
    };
  }, []);

  const applyTheme = (selectedTheme: 'light' | 'dark' | 'system') => {
    if (selectedTheme === 'system') {
      const systemPrefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
      if (systemPrefersDark) {
        document.documentElement.classList.add('dark');
      } else {
        document.documentElement.classList.remove('dark');
      }
    } else if (selectedTheme === 'dark') {
      document.documentElement.classList.add('dark');
    } else {
      document.documentElement.classList.remove('dark');
    }
  };

  const handleThemeChange = (selectedTheme: 'light' | 'dark' | 'system') => {
    setTheme(selectedTheme);
    localStorage.setItem('theme', selectedTheme);
    applyTheme(selectedTheme);
  };

  const handleEmailNotificationsChange = (enabled: boolean) => {
    setEmailNotifications(enabled);
    localStorage.setItem('emailNotifications', enabled.toString());
  };

  const handlePushNotificationsChange = (enabled: boolean) => {
    setPushNotifications(enabled);
    localStorage.setItem('pushNotifications', enabled.toString());
  };

  const handleSmsNotificationsChange = (enabled: boolean) => {
    setSmsNotifications(enabled);
    localStorage.setItem('smsNotifications', enabled.toString());
  };

  const handleLogout = async () => {
    try {
      await dispatch(logout()).unwrap();
      navigate('/login');
    } catch (error) {
      console.error('Logout error:', error);
    }
  };

  const renderContent = () => {
    switch (activeTab) {
      case 'security':
        return <SecurityPanel onBack={() => setActiveTab('appearance')} />;
      case 'notifications':
        return <NotificationsPanel onBack={() => setActiveTab('appearance')} />;
      case 'appearance':
      default:
        return (
          <div className="flex flex-col gap-8">
            <h2 className="text-2xl font-bold text-[#111618] dark:text-white">Appearance</h2>
            
            {/* Theme Settings */}
            <div className="space-y-6">
              <h3 className="text-lg font-semibold text-gray-700 dark:text-gray-300 border-b border-gray-200 dark:border-gray-700 pb-2">Theme Preferences</h3>
              
              <div className="flex flex-col gap-4">
                <div className="flex items-center justify-between p-4 rounded-lg border border-gray-200 dark:border-gray-700">
                  <div>
                    <p className="font-medium text-gray-900 dark:text-white">Light</p>
                    <p className="text-sm text-gray-500 dark:text-gray-400">Light theme</p>
                  </div>
                  <button 
                    onClick={() => handleThemeChange('light')}
                    className={`relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none ${
                      theme === 'light' ? 'bg-primary' : 'bg-gray-300 dark:bg-gray-600'
                    }`}
                  >
                    <span 
                      className={`inline-block h-4 w-4 transform rounded-full bg-white transition-transform ${
                        theme === 'light' ? 'translate-x-6' : 'translate-x-1'
                      }`}
                    />
                  </button>
                </div>
                
                <div className="flex items-center justify-between p-4 rounded-lg border border-gray-200 dark:border-gray-700">
                  <div>
                    <p className="font-medium text-gray-900 dark:text-white">Dark</p>
                    <p className="text-sm text-gray-500 dark:text-gray-400">Dark theme</p>
                  </div>
                  <button 
                    onClick={() => handleThemeChange('dark')}
                    className={`relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none ${
                      theme === 'dark' ? 'bg-primary' : 'bg-gray-300 dark:bg-gray-600'
                    }`}
                  >
                    <span 
                      className={`inline-block h-4 w-4 transform rounded-full bg-white transition-transform ${
                        theme === 'dark' ? 'translate-x-6' : 'translate-x-1'
                      }`}
                    />
                  </button>
                </div>
                
                <div className="flex items-center justify-between p-4 rounded-lg border border-gray-200 dark:border-gray-700">
                  <div>
                    <p className="font-medium text-gray-900 dark:text-white">System Preference</p>
                    <p className="text-sm text-gray-500 dark:text-gray-400">Use system preference for light/dark theme</p>
                  </div>
                  <button 
                    onClick={() => handleThemeChange('system')}
                    className={`relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none ${
                      theme === 'system' ? 'bg-primary' : 'bg-gray-300 dark:bg-gray-600'
                    }`}
                  >
                    <span 
                      className={`inline-block h-4 w-4 transform rounded-full bg-white transition-transform ${
                        theme === 'system' ? 'translate-x-6' : 'translate-x-1'
                      }`}
                    />
                  </button>
                </div>
              </div>
            </div>
            
            {/* Language Settings */}
            <div className="space-y-6">
              <h3 className="text-lg font-semibold text-gray-700 dark:text-gray-300 border-b border-gray-200 dark:border-gray-700 pb-2">Language & Region</h3>
              
              <div className="flex flex-col gap-4">
                <div className="flex items-center justify-between p-4 rounded-lg border border-gray-200 dark:border-gray-700">
                  <div>
                    <p className="font-medium text-gray-900 dark:text-white">Language</p>
                    <p className="text-sm text-gray-500 dark:text-gray-400">Select your preferred language</p>
                  </div>
                  <select className="form-select w-40 rounded-lg border-gray-300 dark:border-gray-600 bg-transparent focus:border-primary focus:ring-primary">
                    <option>English</option>
                    <option>Spanish</option>
                    <option>French</option>
                    <option>German</option>
                  </select>
                </div>
                
                <div className="flex items-center justify-between p-4 rounded-lg border border-gray-200 dark:border-gray-700">
                  <div>
                    <p className="font-medium text-gray-900 dark:text-white">Currency</p>
                    <p className="text-sm text-gray-500 dark:text-gray-400">Select your preferred currency</p>
                  </div>
                  <select className="form-select w-40 rounded-lg border-gray-300 dark:border-gray-600 bg-transparent focus:border-primary focus:ring-primary">
                    <option>USD - US Dollar</option>
                    <option>EUR - Euro</option>
                    <option>GBP - British Pound</option>
                    <option>JPY - Japanese Yen</option>
                  </select>
                </div>
              </div>
            </div>
            
            {/* Notification Settings */}
            <div className="space-y-6">
              <h3 className="text-lg font-semibold text-gray-700 dark:text-gray-300 border-b border-gray-200 dark:border-gray-700 pb-2">Notifications</h3>
              
              <div className="flex flex-col gap-4">
                <div className="flex items-center justify-between p-4 rounded-lg border border-gray-200 dark:border-gray-700">
                  <div>
                    <p className="font-medium text-gray-900 dark:text-white">Email Notifications</p>
                    <p className="text-sm text-gray-500 dark:text-gray-400">Receive email updates about your account</p>
                  </div>
                  <button 
                    onClick={() => handleEmailNotificationsChange(!emailNotifications)}
                    className={`relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none ${
                      emailNotifications ? 'bg-primary' : 'bg-gray-300 dark:bg-gray-600'
                    }`}
                  >
                    <span 
                      className={`inline-block h-4 w-4 transform rounded-full bg-white transition-transform ${
                        emailNotifications ? 'translate-x-6' : 'translate-x-1'
                      }`}
                    />
                  </button>
                </div>
                
                <div className="flex items-center justify-between p-4 rounded-lg border border-gray-200 dark:border-gray-700">
                  <div>
                    <p className="font-medium text-gray-900 dark:text-white">Push Notifications</p>
                    <p className="text-sm text-gray-500 dark:text-gray-400">Receive push notifications on your device</p>
                  </div>
                  <button 
                    onClick={() => handlePushNotificationsChange(!pushNotifications)}
                    className={`relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none ${
                      pushNotifications ? 'bg-primary' : 'bg-gray-300 dark:bg-gray-600'
                    }`}
                  >
                    <span 
                      className={`inline-block h-4 w-4 transform rounded-full bg-white transition-transform ${
                        pushNotifications ? 'translate-x-6' : 'translate-x-1'
                      }`}
                    />
                  </button>
                </div>
                
                <div className="flex items-center justify-between p-4 rounded-lg border border-gray-200 dark:border-gray-700">
                  <div>
                    <p className="font-medium text-gray-900 dark:text-white">SMS Notifications</p>
                    <p className="text-sm text-gray-500 dark:text-gray-400">Receive SMS updates about your transactions</p>
                  </div>
                  <button 
                    onClick={() => handleSmsNotificationsChange(!smsNotifications)}
                    className={`relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none ${
                      smsNotifications ? 'bg-primary' : 'bg-gray-300 dark:bg-gray-600'
                    }`}
                  >
                    <span 
                      className={`inline-block h-4 w-4 transform rounded-full bg-white transition-transform ${
                        smsNotifications ? 'translate-x-6' : 'translate-x-1'
                      }`}
                    />
                  </button>
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
          <h1 className="text-4xl font-black leading-tight tracking-[-0.033em] text-[#111618] dark:text-white">Settings</h1>
          <p className="text-base font-normal leading-normal text-gray-500 dark:text-gray-400">Manage your account preferences and settings.</p>
        </div>
        
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
          {/* Left Column: Navigation */}
          <div className="lg:col-span-1 flex flex-col gap-6">
            <div className="bg-white dark:bg-background-dark rounded-xl border border-gray-200 dark:border-gray-700 shadow-sm p-6">
              <div className="flex flex-col gap-4">
                <button 
                  className="flex items-center gap-3 w-full p-3 rounded-lg text-left hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors"
                  onClick={() => navigate('/profile')}
                >
                  <span className="material-symbols-outlined text-xl">person</span>
                  <span className="font-medium">Profile</span>
                </button>
                <button 
                  className={`flex items-center gap-3 w-full p-3 rounded-lg text-left transition-colors ${
                    activeTab === 'appearance' 
                      ? 'bg-primary/10 dark:bg-primary/20 text-primary' 
                      : 'hover:bg-gray-100 dark:hover:bg-gray-800'
                  }`}
                  onClick={() => setActiveTab('appearance')}
                >
                  <span className="material-symbols-outlined text-xl">palette</span>
                  <span className="font-medium">Appearance</span>
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
                  onClick={handleLogout}
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

export default Settings;