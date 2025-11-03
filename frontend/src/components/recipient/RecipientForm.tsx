import React, { useState } from 'react';

export interface RecipientFormData {
  name: string;
  bankName: string;
  accountNumber: string;
  sortCode?: string;
  iban?: string;
  swiftCode?: string;
  country: string;
  currency: string;
}

interface RecipientFormProps {
  onSubmit: (recipient: RecipientFormData) => void;
  onCancel: () => void;
  initialData?: RecipientFormData;
}

const RecipientForm: React.FC<RecipientFormProps> = ({ onSubmit, onCancel, initialData }) => {
  const [formData, setFormData] = useState<RecipientFormData>(
    initialData || {
      name: '',
      bankName: '',
      accountNumber: '',
      country: 'United Kingdom',
      currency: 'GBP',
    }
  );

  const [errors, setErrors] = useState<Record<string, string>>({});

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));

    // Clear error when user starts typing
    if (errors[name]) {
      setErrors((prev) => {
        const newErrors = { ...prev };
        delete newErrors[name];
        return newErrors;
      });
    }
  };

  const validateForm = () => {
    const newErrors: Record<string, string> = {};

    if (!formData.name.trim()) {
      newErrors.name = 'Name is required';
    }

    if (!formData.bankName.trim()) {
      newErrors.bankName = 'Bank name is required';
    }

    if (!formData.accountNumber.trim()) {
      newErrors.accountNumber = 'Account number is required';
    }

    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (validateForm()) {
      onSubmit(formData);
    }
  };

  // Map countries to their default currencies
  const getCurrencyForCountry = (country: string) => {
    const countryCurrencyMap: Record<string, string> = {
      'United Kingdom': 'GBP',
      Germany: 'EUR',
      Canada: 'CAD',
      'United States': 'USD',
      Australia: 'AUD',
      Nigeria: 'NGN',
      France: 'EUR',
      Japan: 'JPY',
      China: 'CNY',
      India: 'INR',
      Brazil: 'BRL',
      Mexico: 'MXN',
      'South Africa': 'ZAR',
      Kenya: 'KES',
      Ghana: 'GHS',
    };
    return countryCurrencyMap[country] || 'USD';
  };

  const handleCountryChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const country = e.target.value;
    const currency = getCurrencyForCountry(country);
    setFormData({
      ...formData,
      country,
      currency,
    });
  };

  return (
    <form className="flex flex-col gap-8" onSubmit={handleSubmit}>
      <h2 className="text-2xl font-bold text-[#111618] dark:text-white">
        {initialData ? 'Edit Recipient' : 'Add New Recipient'}
      </h2>

      {/* Personal Details Section */}
      <div className="space-y-6">
        <h3 className="text-lg font-semibold text-gray-700 dark:text-gray-300 border-b border-gray-200 dark:border-gray-700 pb-2">
          Personal Details
        </h3>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label
              className="font-medium text-sm text-gray-600 dark:text-gray-300 mb-2 block"
              htmlFor="name"
            >
              Full Name
            </label>
            <input
              className={`form-input w-full rounded-lg border ${
                errors.name ? 'border-red-500' : 'border-gray-300 dark:border-gray-600'
              } bg-transparent focus:border-primary focus:ring-primary`}
              id="name"
              name="name"
              type="text"
              value={formData.name}
              onChange={handleChange}
              required
            />
            {errors.name && <p className="mt-1 text-sm text-red-500">{errors.name}</p>}
          </div>
        </div>
      </div>

      {/* Bank Details Section */}
      <div className="space-y-6">
        <h3 className="text-lg font-semibold text-gray-700 dark:text-gray-300 border-b border-gray-200 dark:border-gray-700 pb-2">
          Bank Details
        </h3>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label
              className="font-medium text-sm text-gray-600 dark:text-gray-300 mb-2 block"
              htmlFor="bankName"
            >
              Bank Name
            </label>
            <input
              className={`form-input w-full rounded-lg border ${
                errors.bankName ? 'border-red-500' : 'border-gray-300 dark:border-gray-600'
              } bg-transparent focus:border-primary focus:ring-primary`}
              id="bankName"
              name="bankName"
              type="text"
              value={formData.bankName}
              onChange={handleChange}
              required
            />
            {errors.bankName && <p className="mt-1 text-sm text-red-500">{errors.bankName}</p>}
          </div>
          <div>
            <label
              className="font-medium text-sm text-gray-600 dark:text-gray-300 mb-2 block"
              htmlFor="accountNumber"
            >
              Account Number
            </label>
            <input
              className={`form-input w-full rounded-lg border ${
                errors.accountNumber ? 'border-red-500' : 'border-gray-300 dark:border-gray-600'
              } bg-transparent focus:border-primary focus:ring-primary`}
              id="accountNumber"
              name="accountNumber"
              type="text"
              value={formData.accountNumber}
              onChange={handleChange}
              required
            />
            {errors.accountNumber && (
              <p className="mt-1 text-sm text-red-500">{errors.accountNumber}</p>
            )}
          </div>
          <div>
            <label
              className="font-medium text-sm text-gray-600 dark:text-gray-300 mb-2 block"
              htmlFor="country"
            >
              Country
            </label>
            <select
              className="form-select w-full rounded-lg border-gray-300 dark:border-gray-600 bg-transparent focus:border-primary focus:ring-primary"
              id="country"
              name="country"
              value={formData.country}
              onChange={handleCountryChange}
            >
              <option value="United Kingdom">United Kingdom</option>
              <option value="Germany">Germany</option>
              <option value="Canada">Canada</option>
              <option value="United States">United States</option>
              <option value="Australia">Australia</option>
              <option value="Nigeria">Nigeria</option>
              <option value="France">France</option>
              <option value="Japan">Japan</option>
              <option value="China">China</option>
              <option value="India">India</option>
              <option value="Brazil">Brazil</option>
              <option value="Mexico">Mexico</option>
              <option value="South Africa">South Africa</option>
              <option value="Kenya">Kenya</option>
              <option value="Ghana">Ghana</option>
            </select>
          </div>
          <div>
            <label
              className="font-medium text-sm text-gray-600 dark:text-gray-300 mb-2 block"
              htmlFor="currency"
            >
              Currency
            </label>
            <select
              className="form-select w-full rounded-lg border-gray-300 dark:border-gray-600 bg-transparent focus:border-primary focus:ring-primary"
              id="currency"
              name="currency"
              value={formData.currency}
              onChange={handleChange}
            >
              <option value="GBP">GBP - British Pound</option>
              <option value="EUR">EUR - Euro</option>
              <option value="CAD">CAD - Canadian Dollar</option>
              <option value="USD">USD - US Dollar</option>
              <option value="AUD">AUD - Australian Dollar</option>
              <option value="NGN">NGN - Nigerian Naira</option>
              <option value="JPY">JPY - Japanese Yen</option>
              <option value="CNY">CNY - Chinese Yuan</option>
              <option value="INR">INR - Indian Rupee</option>
              <option value="BRL">BRL - Brazilian Real</option>
              <option value="MXN">MXN - Mexican Peso</option>
              <option value="ZAR">ZAR - South African Rand</option>
              <option value="KES">KES - Kenyan Shilling</option>
              <option value="GHS">GHS - Ghanaian Cedi</option>
            </select>
          </div>
          <div>
            <label
              className="font-medium text-sm text-gray-600 dark:text-gray-300 mb-2 block"
              htmlFor="sortCode"
            >
              Sort Code (UK only)
            </label>
            <input
              className="form-input w-full rounded-lg border-gray-300 dark:border-gray-600 bg-transparent focus:border-primary focus:ring-primary"
              id="sortCode"
              name="sortCode"
              type="text"
              value={formData.sortCode || ''}
              onChange={handleChange}
              placeholder="12-34-56"
            />
          </div>
          <div>
            <label
              className="font-medium text-sm text-gray-600 dark:text-gray-300 mb-2 block"
              htmlFor="iban"
            >
              IBAN (International)
            </label>
            <input
              className="form-input w-full rounded-lg border-gray-300 dark:border-gray-600 bg-transparent focus:border-primary focus:ring-primary"
              id="iban"
              name="iban"
              type="text"
              value={formData.iban || ''}
              onChange={handleChange}
              placeholder="DE44500105170445678901"
            />
          </div>
          <div>
            <label
              className="font-medium text-sm text-gray-600 dark:text-gray-300 mb-2 block"
              htmlFor="swiftCode"
            >
              SWIFT/BIC Code
            </label>
            <input
              className="form-input w-full rounded-lg border-gray-300 dark:border-gray-600 bg-transparent focus:border-primary focus:ring-primary"
              id="swiftCode"
              name="swiftCode"
              type="text"
              value={formData.swiftCode || ''}
              onChange={handleChange}
              placeholder="COBADEFFXXX"
            />
          </div>
        </div>
      </div>

      <div className="flex gap-3">
        <button
          type="button"
          className="flex min-w-[84px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-12 px-5 bg-gray-100 dark:bg-gray-700 text-gray-800 dark:text-gray-200 gap-2 text-base font-bold leading-normal tracking-[0.015em] hover:bg-gray-200 dark:hover:bg-gray-600 transition-colors"
          onClick={onCancel}
        >
          <span className="truncate">Cancel</span>
        </button>
        <button
          type="submit"
          className="flex min-w-[84px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-12 px-5 bg-primary text-white gap-2 text-base font-bold leading-normal tracking-[0.015em] hover:bg-primary/90 transition-colors"
        >
          <span className="truncate">{initialData ? 'Save Changes' : 'Save Recipient'}</span>
        </button>
      </div>
    </form>
  );
};

export default RecipientForm;
