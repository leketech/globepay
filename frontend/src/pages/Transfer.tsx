import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { currencyService } from '../services/currency.service';
import { transferService } from '../services/transfer.service';

interface FormData {
  recipientName: string;
  sourceAmount: number;
  sourceCurrency: string;
  destCurrency: string;
  recipientCountry: string;
  recipientBankName: string;
  recipientAccountNo: string;
  recipientSwiftCode: string;
  purpose: string;
}

// Define recipient options interface
interface RecipientOption {
  id: number;
  name: string;
}

// Define country-currency mapping directly in the component
const COUNTRIES_AND_CURRENCIES = [
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

const Transfer: React.FC = () => {
  const navigate = useNavigate();
  const [supportedCurrencies, setSupportedCurrencies] = useState<Record<string, string>>({});
  const [formData, setFormData] = useState<FormData>({
    recipientName: '',
    sourceAmount: 1000,
    sourceCurrency: 'USD',
    destCurrency: 'EUR',
    recipientCountry: '',
    recipientBankName: '',
    recipientAccountNo: '',
    recipientSwiftCode: '',
    purpose: ''
  });
  const [exchangeRate, setExchangeRate] = useState(0);
  const [fee, setFee] = useState(0);
  const [destAmount, setDestAmount] = useState(0);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');
  const [showCardModal, setShowCardModal] = useState(false);
  // Define recipient options
  const [recipientOptions] = useState<RecipientOption[]>([
    { id: 1, name: 'Jane Doe' },
    { id: 2, name: 'John Smith' },
    { id: 3, name: 'Emily Jones' }
  ]);

  useEffect(() => {
    // Load supported currencies on component mount
    loadSupportedCurrencies();
  }, []);

  useEffect(() => {
    if (formData.sourceAmount > 0) {
      calculateExchangeRate();
    }
  }, [formData.sourceCurrency, formData.destCurrency, formData.sourceAmount]);

  const loadSupportedCurrencies = async () => {
    try {
      console.log('Loading supported currencies...');
      const currencies = await currencyService.getSupportedCurrencies();
      console.log('Currencies loaded successfully:', currencies);
      console.log('Number of currencies:', Object.keys(currencies).length);
      
      // Validate currencies data
      if (!currencies || Object.keys(currencies).length === 0) {
        throw new Error('No currencies data received');
      }
      
      setSupportedCurrencies(currencies);
      console.log('Supported currencies state updated');
      
      // Clear any previous errors
      if (error && error.includes('Failed to load currency information')) {
        setError('');
      }
    } catch (err) {
      console.error('Failed to load supported currencies:', err);
      if (err instanceof Error) {
        console.error('Error name:', err.name);
        console.error('Error message:', err.message);
        console.error('Error stack:', err.stack);
      }
      const errorMessage = err instanceof Error ? err.message : String(err);
      setError('Failed to load currency information: ' + errorMessage);
    }
  };

  const calculateExchangeRate = async () => {
    try {
      console.log('Calculating exchange rate...', {
        from: formData.sourceCurrency,
        to: formData.destCurrency,
        amount: formData.sourceAmount
      });
      
      // Check if we have valid data
      if (!formData.sourceCurrency || !formData.destCurrency || formData.sourceAmount <= 0) {
        console.log('Skipping exchange rate calculation due to invalid data');
        // Reset values when data is invalid
        setExchangeRate(0);
        setFee(0);
        setDestAmount(0);
        return;
      }
      
      setLoading(true);
      // Get real exchange rate from backend API
      const result = await transferService.getExchangeRate(
        formData.sourceCurrency,
        formData.destCurrency,
        formData.sourceAmount
      );
      
      console.log('Exchange rate calculation result:', result);
      
      // Validate result
      if (!result || !isFinite(result.rate)) {
        throw new Error('Invalid exchange rate calculation result');
      }
      
      setExchangeRate(result.rate);
      setFee(result.fee);
      // Use the convertedAmount from the backend response
      const calculatedDestAmount = result.convertedAmount;
      console.log('Calculated destination amount:', calculatedDestAmount);
      
      // Validate destination amount
      if (isFinite(calculatedDestAmount)) {
        setDestAmount(calculatedDestAmount);
      } else {
        setDestAmount(0);
      }
      
      // Clear any previous errors
      if (error && error.includes('Failed to fetch exchange rate')) {
        setError('');
      }
    } catch (err) {
      console.error('Failed to fetch exchange rate:', err);
      if (err instanceof Error) {
        console.error('Error name:', err.name);
        console.error('Error message:', err.message);
        console.error('Error stack:', err.stack);
      }
      const errorMessage = err instanceof Error ? err.message : String(err);
      setError('Failed to fetch exchange rate: ' + errorMessage);
      // Set default values to avoid NaN
      setExchangeRate(0);
      setFee(0);
      setDestAmount(0);
    } finally {
      setLoading(false);
    }
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      setLoading(true);
      // Mock transfer creation
      setTimeout(() => {
        navigate('/dashboard');
      }, 1000);
    } catch (err) {
      setError('Failed to create transfer');
    } finally {
      setLoading(false);
    }
  };

  // Get currency for selected country
  const getCurrencyForCountry = (countryName: string) => {
    const country = COUNTRIES_AND_CURRENCIES.find((c: any) => c.country === countryName);
    return country ? country.currency : 'USD';
  };

  // Handle country change
  const handleCountryChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const countryName = e.target.value;
    const currency = getCurrencyForCountry(countryName);
    
    console.log('Country changed:', countryName, 'Currency:', currency);
    
    setFormData({
      ...formData,
      recipientCountry: countryName,
      destCurrency: currency
    });
  };

  return (
    <div className="relative flex min-h-screen w-full flex-col group/design-root">

      {/* Main Content */}
      <main className="w-full max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
        <div className="flex flex-col lg:flex-row gap-12">
          {/* Left: Form */}
          <div className="flex-1 lg:w-2/3">
            <div className="flex flex-col gap-8">
              {/* PageHeading */}
              <div className="flex flex-wrap justify-between gap-3 p-4">
                <p className="text-neutral-dark dark:text-white text-4xl font-black leading-tight tracking-[-0.033em]">Send Money</p>
              </div>
              
              {/* Form Container Card */}
              <div className="bg-white dark:bg-background-dark/50 rounded-xl shadow-sm p-6 md:p-8">
                <form className="space-y-8" onSubmit={handleSubmit}>
                  {/* Recipient Section */}
                  <div>
                    <h3 className="text-neutral-dark dark:text-white text-lg font-bold leading-tight tracking-[-0.015em] pb-4">Who are you sending money to?</h3>
                    <div className="flex flex-col sm:flex-row items-end gap-4">
                      <label className="flex flex-col min-w-40 flex-1">
                        <p className="text-neutral-medium dark:text-neutral-light/70 text-sm font-medium leading-normal pb-2">Recipient</p>
                        <select 
                          className="form-select w-full rounded-lg text-neutral-dark dark:text-white focus:outline-0 focus:ring-2 focus:ring-primary/50 border-neutral-light/80 dark:border-neutral-dark/60 bg-white dark:bg-neutral-dark/40 focus:border-primary h-14 placeholder:text-neutral-medium p-[15px] text-base font-normal leading-normal"
                          value={formData.recipientName}
                          onChange={(e) => setFormData({...formData, recipientName: e.target.value})}
                        >
                          <option value="">Search for an existing recipient</option>
                          {recipientOptions.map(recipient => (
                            <option key={recipient.id} value={recipient.name}>{recipient.name}</option>
                          ))}
                        </select>
                      </label>
                      <button 
                        type="button"
                        className="flex w-full sm:w-auto min-w-[84px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-14 px-4 bg-primary/10 dark:bg-primary/20 text-primary dark:text-primary gap-2 text-sm font-bold leading-normal tracking-[0.015em]"
                        onClick={() => navigate('/recipients')}
                      >
                        <span className="material-symbols-outlined text-xl">add</span>
                        <span className="truncate">Add New Recipient</span>
                      </button>
                    </div>
                  </div>
                  
                  {/* Amount Section */}
                  <div>
                    <h3 className="text-neutral-dark dark:text-white text-lg font-bold leading-tight tracking-[-0.015em] pb-4">How much would you like to send?</h3>
                    <div className="grid grid-cols-1 md:grid-cols-2 gap-4 items-start relative">
                      {/* You Send Input */}
                      <div className="relative">
                        <label className="block text-sm font-medium pb-2 text-neutral-medium dark:text-neutral-light/70" htmlFor="you-send">You Send</label>
                        <input 
                          className="form-input w-full rounded-lg text-2xl font-bold text-neutral-dark dark:text-white focus:outline-0 focus:ring-2 focus:ring-primary/50 border-neutral-light/80 dark:border-neutral-dark/60 bg-white dark:bg-neutral-dark/40 focus:border-primary h-14 pl-4 pr-28" 
                          id="you-send" 
                          type="number" 
                          value={formData.sourceAmount || ''}
                          onChange={(e) => setFormData({...formData, sourceAmount: parseFloat(e.target.value) || 0})}
                        />
                        <div className="absolute inset-y-0 right-0 flex items-center pr-3">
                          <select 
                            className="form-select bg-transparent border-0 font-bold text-neutral-dark dark:text-white focus:ring-0"
                            value={formData.sourceCurrency}
                            onChange={(e) => setFormData({...formData, sourceCurrency: e.target.value})}
                          >
                            {Object.entries(supportedCurrencies).map(([code, name]) => (
                              <option key={code} value={code}>
                                {code === 'USD' && '🇺🇸 '} 
                                {code === 'EUR' && '🇪🇺 '} 
                                {code === 'GBP' && '🇬🇧 '} 
                                {code === 'JPY' && '🇯🇵 '} 
                                {code === 'CAD' && '🇨🇦 '} 
                                {code === 'AUD' && '🇦🇺 '} 
                                {code === 'CHF' && '🇨🇭 '} 
                                {code === 'CNY' && '🇨🇳 '} 
                                {code === 'INR' && '🇮🇳 '} 
                                {code === 'NGN' && '🇳🇬 '} 
                                {code} {name !== code ? `(${name})` : ''}
                              </option>
                            ))}
                          </select>
                        </div>
                      </div>
                      
                      {/* Swap Icon */}
                      <div className="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 mt-4 hidden md:block">
                        <button 
                          className="flex h-10 w-10 items-center justify-center rounded-full border-2 border-background-light dark:border-background-dark bg-white dark:bg-neutral-dark/50 text-neutral-medium hover:text-primary" 
                          type="button"
                          onClick={() => {
                            // Swap currencies
                            const tempCurrency = formData.sourceCurrency;
                            setFormData({
                              ...formData,
                              sourceCurrency: formData.destCurrency,
                              destCurrency: tempCurrency
                            });
                          }}
                        >
                          <span className="material-symbols-outlined text-2xl">swap_horiz</span>
                        </button>
                      </div>
                      
                      {/* They Receive Input */}
                      <div className="relative">
                        <label className="block text-sm font-medium pb-2 text-neutral-medium dark:text-neutral-light/70" htmlFor="they-receive">They Receive</label>
                        <input 
                          className="form-input w-full rounded-lg text-2xl font-bold text-neutral-dark dark:text-white focus:outline-0 focus:ring-2 focus:ring-primary/50 border-neutral-light/80 dark:border-neutral-dark/60 bg-white dark:bg-neutral-dark/40 focus:border-primary h-14 pl-4 pr-28" 
                          id="they-receive" 
                          type="text" 
                          value={destAmount.toFixed(2)}
                          readOnly
                        />
                        <div className="absolute inset-y-0 right-0 flex items-center pr-3">
                          <select 
                            className="form-select bg-transparent border-0 font-bold text-neutral-dark dark:text-white focus:ring-0"
                            value={formData.destCurrency}
                            onChange={(e) => setFormData({...formData, destCurrency: e.target.value})}
                          >
                            {Object.entries(supportedCurrencies).map(([code, name]) => (
                              <option key={code} value={code}>
                                {code === 'USD' && '🇺🇸 '} 
                                {code === 'EUR' && '🇪🇺 '} 
                                {code === 'GBP' && '🇬🇧 '} 
                                {code === 'JPY' && '🇯🇵 '} 
                                {code === 'CAD' && '🇨🇦 '} 
                                {code === 'AUD' && '🇦🇺 '} 
                                {code === 'CHF' && '🇨🇭 '} 
                                {code === 'CNY' && '🇨🇳 '} 
                                {code === 'INR' && '🇮🇳 '} 
                                {code === 'NGN' && '🇳🇬 '} 
                                {code} {name !== code ? `(${name})` : ''}
                              </option>
                            ))}
                          </select>
                        </div>
                      </div>
                    </div>
                    
                    {/* Exchange Rate Info */}
                    {exchangeRate > 0 && (
                      <div className="mt-4 p-4 bg-blue-50 dark:bg-blue-900/20 rounded-lg">
                        <div className="flex justify-between items-center">
                          <span className="text-sm text-blue-800 dark:text-blue-200">
                            Exchange Rate: 1 {formData.sourceCurrency} = {exchangeRate.toFixed(4)} {formData.destCurrency}
                          </span>
                          <span className="text-sm text-blue-800 dark:text-blue-200">
                            Fee: {fee.toFixed(2)} {formData.sourceCurrency}
                          </span>
                        </div>
                      </div>
                    )}
                  </div>
                  
                  {/* Payment Method - Add Money via Card */}
                  <div>
                    <h3 className="text-neutral-dark dark:text-white text-lg font-bold leading-tight tracking-[-0.015em] pb-4">Payment Method</h3>
                    <div className="grid grid-cols-1 gap-4">
                      <div className="border border-neutral-light/80 dark:border-neutral-dark/60 rounded-xl p-4">
                        <div className="flex items-center justify-between">
                          <div className="flex items-center">
                            <span className="material-symbols-outlined text-2xl text-primary mr-3">credit_card</span>
                            <div>
                              <p className="font-medium text-neutral-dark dark:text-white">Add Money via Card</p>
                              <p className="text-sm text-neutral-medium dark:text-neutral-light/70">Visa, Mastercard, American Express</p>
                            </div>
                          </div>
                          <button 
                            type="button"
                            className="text-primary font-medium hover:underline"
                            onClick={() => setShowCardModal(true)}
                          >
                            Add Card
                          </button>
                        </div>
                      </div>
                      
                      <div className="border border-neutral-light/80 dark:border-neutral-dark/60 rounded-xl p-4">
                        <div className="flex items-center">
                          <span className="material-symbols-outlined text-2xl text-primary mr-3">account_balance</span>
                          <div>
                            <p className="font-medium text-neutral-dark dark:text-white">Bank Transfer</p>
                            <p className="text-sm text-neutral-medium dark:text-neutral-light/70">Link your bank account</p>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                  
                  {/* Card Modal */}
                  {showCardModal && (
                    <div className="fixed inset-0 bg-black/50 flex items-center justify-center z-50 p-4">
                      <div className="bg-white dark:bg-background-dark rounded-xl shadow-lg w-full max-w-md">
                        <div className="p-6">
                          <div className="flex justify-between items-center mb-4">
                            <h3 className="text-xl font-bold text-neutral-dark dark:text-white">Add Card</h3>
                            <button 
                              type="button"
                              className="text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200"
                              onClick={() => setShowCardModal(false)}
                            >
                              <span className="material-symbols-outlined">close</span>
                            </button>
                          </div>
                          
                          <form className="space-y-4">
                            <div>
                              <label className="block text-sm font-medium text-neutral-medium dark:text-neutral-light/70 mb-2">Card Number</label>
                              <input 
                                type="text" 
                                placeholder="1234 5678 9012 3456"
                                className="w-full rounded-lg border-neutral-light/80 dark:border-neutral-dark/60 bg-white dark:bg-neutral-dark/40 focus:border-primary focus:ring-primary p-3"
                              />
                            </div>
                            
                            <div className="grid grid-cols-2 gap-4">
                              <div>
                                <label className="block text-sm font-medium text-neutral-medium dark:text-neutral-light/70 mb-2">Expiry Date</label>
                                <input 
                                  type="text" 
                                  placeholder="MM/YY"
                                  className="w-full rounded-lg border-neutral-light/80 dark:border-neutral-dark/60 bg-white dark:bg-neutral-dark/40 focus:border-primary focus:ring-primary p-3"
                                />
                              </div>
                              
                              <div>
                                <label className="block text-sm font-medium text-neutral-medium dark:text-neutral-light/70 mb-2">CVV</label>
                                <input 
                                  type="text" 
                                  placeholder="123"
                                  className="w-full rounded-lg border-neutral-light/80 dark:border-neutral-dark/60 bg-white dark:bg-neutral-dark/40 focus:border-primary focus:ring-primary p-3"
                                />
                              </div>
                            </div>
                            
                            <div>
                              <label className="block text-sm font-medium text-neutral-medium dark:text-neutral-light/70 mb-2">Cardholder Name</label>
                              <input 
                                type="text" 
                                placeholder="John Doe"
                                className="w-full rounded-lg border-neutral-light/80 dark:border-neutral-dark/60 bg-white dark:bg-neutral-dark/40 focus:border-primary focus:ring-primary p-3"
                              />
                            </div>
                            
                            <div className="flex justify-end gap-3 pt-4">
                              <button 
                                type="button"
                                className="px-4 py-2 text-neutral-medium dark:text-neutral-light/70 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-lg"
                                onClick={() => setShowCardModal(false)}
                              >
                                Cancel
                              </button>
                              <button 
                                type="button"
                                className="px-4 py-2 bg-primary text-white rounded-lg hover:bg-primary/90"
                                onClick={() => {
                                  // In a real implementation, this would save the card
                                  setShowCardModal(false);
                                  // Show success message
                                  alert('Card added successfully!');
                                }}
                              >
                                Add Card
                              </button>
                            </div>
                          </form>
                        </div>
                      </div>
                    </div>
                  )}
                  
                  {/* Recipient Country */}
                  <div>
                    <h3 className="text-neutral-dark dark:text-white text-lg font-bold leading-tight tracking-[-0.015em] pb-4">Recipient Country</h3>
                    <label className="flex flex-col min-w-40 flex-1">
                      <select 
                        className="form-select w-full rounded-lg text-neutral-dark dark:text-white focus:outline-0 focus:ring-2 focus:ring-primary/50 border-neutral-light/80 dark:border-neutral-dark/60 bg-white dark:bg-neutral-dark/40 focus:border-primary h-14 placeholder:text-neutral-medium p-[15px] text-base font-normal leading-normal"
                        name="recipientCountry"
                        value={formData.recipientCountry}
                        onChange={handleCountryChange}
                      >
                        <option value="">Select recipient country</option>
                        {COUNTRIES_AND_CURRENCIES.map((country) => (
                          <option key={country.code} value={country.country}>
                            {country.country}
                          </option>
                        ))}
                      </select>
                    </label>
                  </div>
                  
                  {/* Purpose */}
                  <div>
                    <h3 className="text-neutral-dark dark:text-white text-lg font-bold leading-tight tracking-[-0.015em] pb-4">Purpose of Transfer</h3>
                    <label className="flex flex-col min-w-40 flex-1">
                      <select 
                        className="form-select w-full rounded-lg text-neutral-dark dark:text-white focus:outline-0 focus:ring-2 focus:ring-primary/50 border-neutral-light/80 dark:border-neutral-dark/60 bg-white dark:bg-neutral-dark/40 focus:border-primary h-14 placeholder:text-neutral-medium p-[15px] text-base font-normal leading-normal"
                        name="purpose"
                        value={formData.purpose}
                        onChange={handleChange}
                      >
                        <option value="">Select purpose</option>
                        <option value="family_support">Family Support</option>
                        <option value="education">Education</option>
                        <option value="medical">Medical Treatment</option>
                        <option value="business">Business Investment</option>
                        <option value="travel">Travel & Tourism</option>
                        <option value="gift">Gift</option>
                        <option value="other">Other</option>
                      </select>
                    </label>
                  </div>
                  
                  {/* Submit Button */}
                  <button 
                    type="submit"
                    disabled={loading || !formData.recipientName || !formData.sourceAmount || !formData.recipientCountry}
                    className="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-14 px-6 bg-primary text-white text-base font-bold leading-normal tracking-[0.015em] hover:bg-primary/90 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
                  >
                    {loading ? (
                      <span className="flex items-center">
                        <span className="material-symbols-outlined animate-spin mr-2">progress_activity</span>
                        Processing...
                      </span>
                    ) : (
                      <span className="truncate">Continue</span>
                    )}
                  </button>
                  
                  {error && (
                    <div className="mt-4 p-4 bg-red-50 dark:bg-red-900/20 text-red-700 dark:text-red-200 rounded-lg">
                      {error}
                    </div>
                  )}
                </form>
              </div>
            </div>
          </div>
          
          {/* Right: Summary Panel */}
          <div className="w-full lg:w-1/3">
            <div className="bg-white dark:bg-background-dark/50 rounded-xl shadow-sm p-6 sticky top-6">
              <h3 className="text-neutral-dark dark:text-white text-lg font-bold leading-tight tracking-[-0.015em] pb-4">Transfer Summary</h3>
              
              <div className="space-y-4">
                <div className="flex justify-between">
                  <span className="text-neutral-medium dark:text-neutral-light/70">You send</span>
                  <span className="font-medium text-neutral-dark dark:text-white">
                    {formData.sourceAmount.toFixed(2)} {formData.sourceCurrency}
                  </span>
                </div>
                
                <div className="flex justify-between">
                  <span className="text-neutral-medium dark:text-neutral-light/70">Exchange rate</span>
                  <span className="font-medium text-neutral-dark dark:text-white">
                    {exchangeRate > 0 ? exchangeRate.toFixed(4) : '—'}
                  </span>
                </div>
                
                <div className="flex justify-between">
                  <span className="text-neutral-medium dark:text-neutral-light/70">Fee</span>
                  <span className="font-medium text-neutral-dark dark:text-white">
                    {fee.toFixed(2)} {formData.sourceCurrency}
                  </span>
                </div>
                
                <div className="border-t border-neutral-light/80 dark:border-neutral-dark/60 pt-4">
                  <div className="flex justify-between">
                    <span className="text-neutral-medium dark:text-neutral-light/70">Total to pay</span>
                    <span className="font-bold text-neutral-dark dark:text-white">
                      {(formData.sourceAmount + fee).toFixed(2)} {formData.sourceCurrency}
                    </span>
                  </div>
                </div>
                
                <div className="border-t border-neutral-light/80 dark:border-neutral-dark/60 pt-4">
                  <div className="flex justify-between">
                    <span className="text-neutral-medium dark:text-neutral-light/70">Recipient gets</span>
                    <span className="font-bold text-neutral-dark dark:text-white">
                      {isFinite(destAmount) ? destAmount.toFixed(2) : '0.00'} {formData.destCurrency}
                    </span>
                  </div>
                </div>
                
                <div className="pt-4">
                  <div className="flex items-center text-sm text-neutral-medium dark:text-neutral-light/70">
                    <span className="material-symbols-outlined text-base mr-2">info</span>
                    <span>Estimated arrival: 1-2 business days</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>
  );
};

export default Transfer;