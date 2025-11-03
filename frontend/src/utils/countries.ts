export interface CountryCurrency {
  code: string;
  country: string;
  currency: string;
}

export const COUNTRIES_AND_CURRENCIES: CountryCurrency[] = [
  { code: 'US', country: 'United States', currency: 'USD' },
  { code: 'GB', country: 'United Kingdom', currency: 'GBP' },
  { code: 'EU', country: 'European Union', currency: 'EUR' },
  { code: 'NG', country: 'Nigeria', currency: 'NGN' },
  { code: 'CA', country: 'Canada', currency: 'CAD' },
  { code: 'AU', country: 'Australia', currency: 'AUD' },
  { code: 'JP', country: 'Japan', currency: 'JPY' },
  { code: 'CN', country: 'China', currency: 'CNY' },
  { code: 'IN', country: 'India', currency: 'INR' },
  { code: 'BR', country: 'Brazil', currency: 'BRL' },
  { code: 'MX', country: 'Mexico', currency: 'MXN' },
  { code: 'ZA', country: 'South Africa', currency: 'ZAR' },
  { code: 'KE', country: 'Kenya', currency: 'KES' },
  { code: 'GH', country: 'Ghana', currency: 'GHS' },
  { code: 'FR', country: 'France', currency: 'EUR' },
  { code: 'DE', country: 'Germany', currency: 'EUR' },
];
