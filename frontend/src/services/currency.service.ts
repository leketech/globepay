// Currency exchange service using ExchangeRate-API.com
const EXCHANGE_RATE_API_BASE = 'https://api.exchangerate-api.com/v4';

export interface ExchangeRate {
  success: boolean;
  base: string;
  date: string;
  rates: Record<string, number>;
}

export interface CurrencyConversion {
  success: boolean;
  from: string;
  to: string;
  amount: number;
  rate: number;
  result: number;
}

export interface SupportedCurrencies {
  success: boolean;
  rates: Record<string, number>;
}

export const currencyService = {
  /**
   * Get latest exchange rates for a base currency
   * @param baseCurrency The base currency (e.g., 'USD')
   * @returns ExchangeRate object with rates
   */
  async getLatestRates(baseCurrency = 'USD'): Promise<ExchangeRate> {
    try {
      const url = `${EXCHANGE_RATE_API_BASE}/latest/${baseCurrency}`;
      console.log('Fetching exchange rates from:', url);

      const controller = new AbortController();
      const timeoutId = setTimeout(() => controller.abort(), 10000); // 10 second timeout

      const response = await fetch(url, {
        signal: controller.signal,
      });
      clearTimeout(timeoutId);

      if (!response.ok) {
        const errorText = await response.text();
        console.error(`HTTP error! status: ${response.status}, body: ${errorText}`);
        throw new Error(`HTTP error! status: ${response.status}, body: ${errorText}`);
      }

      const data = await response.json();
      console.log('Exchange rates response:', data);

      // Validate response structure
      if (!data || !data.base || !data.rates) {
        throw new Error('Invalid response structure from currency API');
      }

      return {
        success: true,
        base: data.base,
        date: data.date,
        rates: data.rates,
      };
    } catch (error) {
      console.error('Failed to fetch exchange rates:', error);
      if (error instanceof Error && error.name === 'AbortError') {
        throw new Error('Request timeout: The currency API is not responding');
      }
      throw error;
    }
  },

  /**
   * Convert an amount from one currency to another
   * @param fromCurrency The source currency (e.g., 'USD')
   * @param toCurrency The target currency (e.g., 'EUR')
   * @param amount The amount to convert
   * @returns CurrencyConversion object with result
   */
  async convertCurrency(
    fromCurrency: string,
    toCurrency: string,
    amount: number
  ): Promise<CurrencyConversion> {
    try {
      console.log('Converting currency...', { fromCurrency, toCurrency, amount });
      const rates = await this.getLatestRates(fromCurrency);
      console.log('Rates received:', rates);

      // Validate that the target currency exists
      if (!(toCurrency in rates.rates)) {
        console.error(
          `Currency ${toCurrency} not found in rates. Available currencies:`,
          Object.keys(rates.rates)
        );
        throw new Error(`Exchange rate not available for ${fromCurrency} to ${toCurrency}`);
      }

      const rate = rates.rates[toCurrency];
      console.log('Rate for', toCurrency, ':', rate);

      if (!rate || typeof rate !== 'number') {
        throw new Error(`Invalid exchange rate for ${fromCurrency} to ${toCurrency}: ${rate}`);
      }

      const result = amount * rate;
      console.log('Conversion result:', { result, rate, amount });

      // Validate result
      if (!isFinite(result)) {
        throw new Error(`Invalid conversion result: ${result}`);
      }

      return {
        success: true,
        from: fromCurrency,
        to: toCurrency,
        amount: amount,
        rate: rate,
        result: result,
      };
    } catch (error) {
      console.error('Failed to convert currency:', error);
      throw error;
    }
  },

  /**
   * Get historical exchange rates for a specific date
   * Note: ExchangeRate-API.com free tier doesn't support historical rates
   * @returns Promise that rejects with an error
   */
  async getHistoricalRates(): Promise<ExchangeRate> {
    throw new Error('Historical rates not available in free tier');
  },

  /**
   * Get a list of supported currencies
   * @returns Object with currency codes as keys and names as values
   */
  async getSupportedCurrencies(): Promise<Record<string, string>> {
    try {
      console.log('Fetching supported currencies...');
      // Get rates for USD as base currency to get list of supported currencies
      const ratesData = await this.getLatestRates('USD');
      console.log('Rates data received:', ratesData);

      // Validate response
      if (!ratesData || !ratesData.rates) {
        throw new Error('Invalid rates data received from API');
      }

      // Map currency codes to full names
      // Note: ExchangeRate-API.com doesn't provide full names, so we'll use codes as names
      const currencies: Record<string, string> = {};
      currencies['USD'] = 'US Dollar';
      currencies['EUR'] = 'Euro';
      currencies['GBP'] = 'British Pound';
      currencies['JPY'] = 'Japanese Yen';
      currencies['CAD'] = 'Canadian Dollar';
      currencies['AUD'] = 'Australian Dollar';
      currencies['CHF'] = 'Swiss Franc';
      currencies['CNY'] = 'Chinese Yuan';
      currencies['INR'] = 'Indian Rupee';
      currencies['NGN'] = 'Nigerian Naira';

      // Add other currencies from the rates response
      Object.keys(ratesData.rates).forEach((code) => {
        if (!currencies[code]) {
          currencies[code] = code; // Use code as name if not in our predefined list
        }
      });

      console.log('Supported currencies:', currencies);
      return currencies;
    } catch (error) {
      console.error('Failed to fetch supported currencies:', error);
      throw error;
    }
  },
};
