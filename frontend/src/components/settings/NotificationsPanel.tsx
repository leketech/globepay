import React, { useState, useEffect } from 'react';
import { userPreferencesService, UserPreferences } from '../../services/userPreferences.service';

interface NotificationsPanelProps {
  onBack: () => void;
}

const NotificationsPanel: React.FC<NotificationsPanelProps> = ({ onBack }) => {
  const [emailNotifications, setEmailNotifications] = useState(true);
  const [pushNotifications, setPushNotifications] = useState(false);
  const [smsNotifications, setSmsNotifications] = useState(false);
  const [transactionAlerts, setTransactionAlerts] = useState(true);
  const [securityAlerts, setSecurityAlerts] = useState(true);
  const [marketingEmails, setMarketingEmails] = useState(false);

  // Load saved preferences from backend on component mount
  useEffect(() => {
    const loadPreferences = async () => {
      try {
        const token = localStorage.getItem('token') || '';
        if (token) {
          const preferences = await userPreferencesService.getUserPreferences(token);
          setEmailNotifications(preferences.email_notifications);
          setPushNotifications(preferences.push_notifications);
          setSmsNotifications(preferences.sms_notifications);
          setTransactionAlerts(preferences.transaction_alerts);
          setSecurityAlerts(preferences.security_alerts);
          setMarketingEmails(preferences.marketing_emails);
        }
      } catch (error) {
        console.error('Failed to load preferences:', error);
      }
    };

    loadPreferences();
  }, []);

  const savePreference = async (updates: Partial<UserPreferences>) => {
    try {
      const token = localStorage.getItem('token') || '';
      if (token) {
        await userPreferencesService.updateUserPreferences(token, updates);
      }
    } catch (error) {
      console.error('Failed to save preferences:', error);
    }
  };

  const handleEmailNotificationsChange = (enabled: boolean) => {
    setEmailNotifications(enabled);
    savePreference({ email_notifications: enabled });
  };

  const handlePushNotificationsChange = (enabled: boolean) => {
    setPushNotifications(enabled);
    savePreference({ push_notifications: enabled });
  };

  const handleSmsNotificationsChange = (enabled: boolean) => {
    setSmsNotifications(enabled);
    savePreference({ sms_notifications: enabled });
  };

  const handleTransactionAlertsChange = (enabled: boolean) => {
    setTransactionAlerts(enabled);
    savePreference({ transaction_alerts: enabled });
  };

  const handleSecurityAlertsChange = (enabled: boolean) => {
    setSecurityAlerts(enabled);
    savePreference({ security_alerts: enabled });
  };

  const handleMarketingEmailsChange = (enabled: boolean) => {
    setMarketingEmails(enabled);
    savePreference({ marketing_emails: enabled });
  };

  return (
    <div className="flex flex-col gap-8">
      <div className="flex items-center gap-4">
        <button 
          onClick={onBack}
          className="flex items-center gap-2 text-primary hover:underline"
        >
          <span className="material-symbols-outlined">arrow_back</span>
          <span>Back to Settings</span>
        </button>
      </div>
      
      <h2 className="text-2xl font-bold text-[#111618] dark:text-white">Notifications</h2>
      
      {/* Notification Preferences */}
      <div className="space-y-6">
        <h3 className="text-lg font-semibold text-gray-700 dark:text-gray-300 border-b border-gray-200 dark:border-gray-700 pb-2">Notification Preferences</h3>
        
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
      
      {/* Transaction Notifications */}
      <div className="space-y-6">
        <h3 className="text-lg font-semibold text-gray-700 dark:text-gray-300 border-b border-gray-200 dark:border-gray-700 pb-2">Transaction Notifications</h3>
        
        <div className="flex flex-col gap-4">
          <div className="flex items-center justify-between p-4 rounded-lg border border-gray-200 dark:border-gray-700">
            <div>
              <p className="font-medium text-gray-900 dark:text-white">Transaction Alerts</p>
              <p className="text-sm text-gray-500 dark:text-gray-400">Get notified when transactions occur</p>
            </div>
            <button 
              onClick={() => handleTransactionAlertsChange(!transactionAlerts)}
              className={`relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none ${
                transactionAlerts ? 'bg-primary' : 'bg-gray-300 dark:bg-gray-600'
              }`}
            >
              <span 
                className={`inline-block h-4 w-4 transform rounded-full bg-white transition-transform ${
                  transactionAlerts ? 'translate-x-6' : 'translate-x-1'
                }`}
              />
            </button>
          </div>
          
          <div className="flex items-center justify-between p-4 rounded-lg border border-gray-200 dark:border-gray-700">
            <div>
              <p className="font-medium text-gray-900 dark:text-white">Security Alerts</p>
              <p className="text-sm text-gray-500 dark:text-gray-400">Get notified about security-related events</p>
            </div>
            <button 
              onClick={() => handleSecurityAlertsChange(!securityAlerts)}
              className={`relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none ${
                securityAlerts ? 'bg-primary' : 'bg-gray-300 dark:bg-gray-600'
              }`}
            >
              <span 
                className={`inline-block h-4 w-4 transform rounded-full bg-white transition-transform ${
                  securityAlerts ? 'translate-x-6' : 'translate-x-1'
                }`}
              />
            </button>
          </div>
        </div>
      </div>
      
      {/* Marketing Communications */}
      <div className="space-y-6">
        <h3 className="text-lg font-semibold text-gray-700 dark:text-gray-300 border-b border-gray-200 dark:border-gray-700 pb-2">Marketing Communications</h3>
        
        <div className="flex flex-col gap-4">
          <div className="flex items-center justify-between p-4 rounded-lg border border-gray-200 dark:border-gray-700">
            <div>
              <p className="font-medium text-gray-900 dark:text-white">Marketing Emails</p>
              <p className="text-sm text-gray-500 dark:text-gray-400">Receive promotional emails and updates</p>
            </div>
            <button 
              onClick={() => handleMarketingEmailsChange(!marketingEmails)}
              className={`relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none ${
                marketingEmails ? 'bg-primary' : 'bg-gray-300 dark:bg-gray-600'
              }`}
            >
              <span 
                className={`inline-block h-4 w-4 transform rounded-full bg-white transition-transform ${
                  marketingEmails ? 'translate-x-6' : 'translate-x-1'
                }`}
              />
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default NotificationsPanel;