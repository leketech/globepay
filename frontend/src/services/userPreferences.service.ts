import { userApi } from './api';

export interface UserPreferences {
  id: string;
  user_id: string;
  email_notifications: boolean;
  push_notifications: boolean;
  sms_notifications: boolean;
  transaction_alerts: boolean;
  security_alerts: boolean;
  marketing_emails: boolean;
  two_factor_enabled: boolean;
  created_at: string;
  updated_at: string;
}

export const userPreferencesService = {
  async getUserPreferences(token: string): Promise<UserPreferences> {
    try {
      const response = await userApi.getUserPreferences(token);
      return response;
    } catch (error) {
      console.error('Failed to fetch user preferences:', error);
      // Return default preferences if fetch fails
      return {
        id: '',
        user_id: '',
        email_notifications: true,
        push_notifications: false,
        sms_notifications: false,
        transaction_alerts: true,
        security_alerts: true,
        marketing_emails: false,
        two_factor_enabled: false,
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString(),
      };
    }
  },

  async updateUserPreferences(
    token: string,
    preferences: Partial<UserPreferences>
  ): Promise<UserPreferences> {
    // Map our preferences to the API's expected format
    const apiPreferences: { [key: string]: any } = {};

    // Map boolean fields directly
    if (preferences.email_notifications !== undefined) {
      apiPreferences.email_notifications = preferences.email_notifications;
    }
    if (preferences.push_notifications !== undefined) {
      apiPreferences.push_notifications = preferences.push_notifications;
    }
    if (preferences.sms_notifications !== undefined) {
      apiPreferences.sms_notifications = preferences.sms_notifications;
    }
    if (preferences.transaction_alerts !== undefined) {
      apiPreferences.transaction_alerts = preferences.transaction_alerts;
    }
    if (preferences.security_alerts !== undefined) {
      apiPreferences.security_alerts = preferences.security_alerts;
    }
    if (preferences.marketing_emails !== undefined) {
      apiPreferences.marketing_emails = preferences.marketing_emails;
    }
    if (preferences.two_factor_enabled !== undefined) {
      apiPreferences.two_factor_enabled = preferences.two_factor_enabled;
    }

    try {
      const response = await userApi.updateUserPreferences(token, apiPreferences);
      return response;
    } catch (error) {
      console.error('Failed to update user preferences:', error);
      throw error;
    }
  },
};
