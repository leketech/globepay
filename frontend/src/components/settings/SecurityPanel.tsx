import React, { useState, useEffect } from 'react';
import { userPreferencesService } from '../../services/userPreferences.service';

interface SecurityPanelProps {
  onBack: () => void;
}

const SecurityPanel: React.FC<SecurityPanelProps> = ({ onBack }) => {
  const [currentPassword, setCurrentPassword] = useState('');
  const [newPassword, setNewPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [twoFactorEnabled, setTwoFactorEnabled] = useState(false);
  const [loading, setLoading] = useState(false);

  // Load saved preferences from backend on component mount
  useEffect(() => {
    const loadPreferences = async () => {
      try {
        const token = localStorage.getItem('token') || '';
        if (token) {
          const preferences = await userPreferencesService.getUserPreferences(token);
          setTwoFactorEnabled(preferences.two_factor_enabled);
        }
      } catch (error) {
        console.error('Failed to load preferences:', error);
      }
    };

    loadPreferences();
  }, []);

  const handlePasswordChange = async (e: React.FormEvent) => {
    e.preventDefault();
    if (newPassword !== confirmPassword) {
      alert('New passwords do not match');
      return;
    }

    setLoading(true);
    try {
      // In a real app, this would call an API to update the password
      console.log('Password change requested');
      // Example: await api.changePassword(currentPassword, newPassword);
      alert('Password updated successfully!');

      // Reset form
      setCurrentPassword('');
      setNewPassword('');
      setConfirmPassword('');
    } catch (error) {
      alert('Failed to update password. Please try again.');
      console.error('Password update error:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleTwoFactorChange = async (enabled: boolean) => {
    setTwoFactorEnabled(enabled);
    try {
      const token = localStorage.getItem('token') || '';
      if (token) {
        await userPreferencesService.updateUserPreferences(token, { two_factor_enabled: enabled });
        alert(enabled ? 'Two-factor authentication enabled' : 'Two-factor authentication disabled');
      }
    } catch (error) {
      console.error('Failed to update two-factor setting:', error);
      alert('Failed to update two-factor authentication setting. Please try again.');
    }
  };

  return (
    <div className="flex flex-col gap-8">
      <div className="flex items-center gap-4">
        <button onClick={onBack} className="flex items-center gap-2 text-primary hover:underline">
          <span className="material-symbols-outlined">arrow_back</span>
          <span>Back to Settings</span>
        </button>
      </div>

      <h2 className="text-2xl font-bold text-[#111618] dark:text-white">Security</h2>

      {/* Password Change */}
      <div className="space-y-6">
        <h3 className="text-lg font-semibold text-gray-700 dark:text-gray-300 border-b border-gray-200 dark:border-gray-700 pb-2">
          Change Password
        </h3>

        <form onSubmit={handlePasswordChange} className="space-y-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
              Current Password
            </label>
            <input
              type="password"
              value={currentPassword}
              onChange={(e) => setCurrentPassword(e.target.value)}
              className="w-full rounded-lg border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-800 focus:border-primary focus:ring-primary"
              required
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
              New Password
            </label>
            <input
              type="password"
              value={newPassword}
              onChange={(e) => setNewPassword(e.target.value)}
              className="w-full rounded-lg border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-800 focus:border-primary focus:ring-primary"
              required
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
              Confirm New Password
            </label>
            <input
              type="password"
              value={confirmPassword}
              onChange={(e) => setConfirmPassword(e.target.value)}
              className="w-full rounded-lg border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-800 focus:border-primary focus:ring-primary"
              required
            />
          </div>

          <button
            type="submit"
            disabled={loading}
            className="px-4 py-2 bg-primary text-white rounded-lg hover:bg-primary/90 transition-colors disabled:opacity-50"
          >
            {loading ? 'Updating...' : 'Update Password'}
          </button>
        </form>
      </div>

      {/* Two-Factor Authentication */}
      <div className="space-y-6">
        <h3 className="text-lg font-semibold text-gray-700 dark:text-gray-300 border-b border-gray-200 dark:border-gray-700 pb-2">
          Two-Factor Authentication
        </h3>

        <div className="flex items-center justify-between p-4 rounded-lg border border-gray-200 dark:border-gray-700">
          <div>
            <p className="font-medium text-gray-900 dark:text-white">Two-Factor Authentication</p>
            <p className="text-sm text-gray-500 dark:text-gray-400">
              Add an extra layer of security to your account
            </p>
          </div>
          <button
            onClick={() => handleTwoFactorChange(!twoFactorEnabled)}
            className={`relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none ${
              twoFactorEnabled ? 'bg-primary' : 'bg-gray-300 dark:bg-gray-600'
            }`}
          >
            <span
              className={`inline-block h-4 w-4 transform rounded-full bg-white transition-transform ${
                twoFactorEnabled ? 'translate-x-6' : 'translate-x-1'
              }`}
            />
          </button>
        </div>
      </div>

      {/* Login Activity */}
      <div className="space-y-6">
        <h3 className="text-lg font-semibold text-gray-700 dark:text-gray-300 border-b border-gray-200 dark:border-gray-700 pb-2">
          Recent Login Activity
        </h3>

        <div className="space-y-4">
          <div className="p-4 rounded-lg border border-gray-200 dark:border-gray-700">
            <div className="flex justify-between">
              <div>
                <p className="font-medium text-gray-900 dark:text-white">Chrome on Windows</p>
                <p className="text-sm text-gray-500 dark:text-gray-400">
                  New York, US • Today at 10:30 AM
                </p>
              </div>
              <span className="text-green-500 text-sm font-medium">Current Session</span>
            </div>
          </div>

          <div className="p-4 rounded-lg border border-gray-200 dark:border-gray-700">
            <div className="flex justify-between">
              <div>
                <p className="font-medium text-gray-900 dark:text-white">Safari on iPhone</p>
                <p className="text-sm text-gray-500 dark:text-gray-400">
                  New York, US • Yesterday at 4:15 PM
                </p>
              </div>
              <span className="text-gray-500 dark:text-gray-400 text-sm">Active</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default SecurityPanel;
