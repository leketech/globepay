import React from 'react';
import { useNavigate } from 'react-router-dom';

const Welcome: React.FC = () => {
  const navigate = useNavigate();

  return (
    <div className="relative flex min-h-screen w-full flex-col group/design-root overflow-x-hidden">
      {/* TopNavBar */}
      <header className="sticky top-0 z-50 w-full bg-background-light/80 dark:bg-background-dark/80 backdrop-blur-sm">
        <div className="flex items-center justify-between whitespace-nowrap border-b border-solid border-border-light dark:border-border-dark px-6 md:px-10 py-3 max-w-7xl mx-auto">
          <div className="flex items-center gap-4">
            <div className="size-8 text-primary">
              <svg fill="none" viewBox="0 0 48 48" xmlns="http://www.w3.org/2000/svg">
                <path
                  d="M42.4379 44C42.4379 44 36.0744 33.9038 41.1692 24C46.8624 12.9336 42.2078 4 42.2078 4L7.01134 4C7.01134 4 11.6577 12.932 5.96912 23.9969C0.876273 33.9029 7.27094 44 7.27094 44L42.4379 44Z"
                  fill="currentColor"
                ></path>
              </svg>
            </div>
            <h2 className="text-xl font-bold leading-tight tracking-[-0.015em]">Globepay</h2>
          </div>
          <nav className="hidden md:flex flex-1 justify-center items-center gap-9">
            <a
              className="text-sm font-medium leading-normal hover:text-primary dark:hover:text-secondary transition-colors"
              href="#"
            >
              Features
            </a>
            <a
              className="text-sm font-medium leading-normal hover:text-primary dark:hover:text-secondary transition-colors"
              href="#"
            >
              How It Works
            </a>
            <a
              className="text-sm font-medium leading-normal hover:text-primary dark:hover:text-secondary transition-colors"
              href="#"
            >
              Help
            </a>
          </nav>
          <div className="flex items-center gap-2">
            <button
              onClick={() => navigate('/login')}
              className="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-10 px-4 bg-neutral-light dark:bg-neutral-dark text-text-light dark:text-text-dark text-sm font-bold leading-normal tracking-[0.015em] hover:bg-neutral-light/80 dark:hover:bg-neutral-dark/80 transition-colors"
            >
              <span className="truncate">Log In</span>
            </button>
            <button
              onClick={() => navigate('/signup')}
              className="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-10 px-4 bg-primary text-white text-sm font-bold leading-normal tracking-[0.015em] hover:bg-primary/90 transition-colors"
            >
              <span className="truncate">Sign Up</span>
            </button>
          </div>
        </div>
      </header>

      <main className="flex-grow">
        {/* HeroSection */}
        <section className="w-full py-16 md:py-24">
          <div className="max-w-7xl mx-auto px-6 md:px-10">
            <div className="grid grid-cols-1 lg:grid-cols-2 gap-12 items-center">
              <div className="flex flex-col gap-6 text-left">
                <h1 className="text-4xl font-black leading-tight tracking-[-0.033em] md:text-5xl lg:text-6xl">
                  Send Money Across the Globe. Fast, Simple, Secure.
                </h1>
                <p className="text-base font-normal leading-normal md:text-lg text-text-light/80 dark:text-text-dark/80">
                  Join thousands who trust Globepay for low-fee international transfers.
                </p>
                <button
                  onClick={() => navigate('/signup')}
                  className="flex min-w-[84px] max-w-[240px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-12 px-5 bg-primary text-white text-base font-bold leading-normal tracking-[0.015em] hover:bg-primary/90 transition-colors"
                >
                  <span className="truncate">Send Money Now</span>
                </button>
              </div>
              <div className="flex items-center justify-center">
                <div className="w-full max-w-md bg-white dark:bg-neutral-dark p-6 rounded-xl shadow-lg">
                  <div className="flex flex-col gap-4">
                    <div className="flex flex-col gap-2">
                      <label className="text-sm font-medium" htmlFor="send-amount">
                        You send
                      </label>
                      <div className="flex items-center border border-border-light dark:border-border-dark rounded-lg overflow-hidden focus-within:ring-2 focus-within:ring-primary">
                        <input
                          className="flex-grow p-3 bg-transparent border-0 focus:ring-0"
                          id="send-amount"
                          type="number"
                          value="1000"
                        />
                        <div className="flex items-center gap-2 p-3 bg-neutral-light dark:bg-background-dark">
                          <img
                            alt="USA Flag"
                            className="h-5 rounded-sm"
                            src="https://lh3.googleusercontent.com/aida-public/AB6AXuA6fXxCZ0pK0Cvx7rx5RsU7Ni8DPzMXwV87IW4jef71HPhz6AXNOh1GPaYiwlFC-r8lZe_r02mP4Yp5-VLEoQcvSPXZwSHDAvG0WbtPcx-zpFEmxB-Ko3yEtDNbfmigIMgLSnikYy3e6R2iG4XzrYhtv4Dvoc54NrteQewhCCo3zTFcXHpln69dg_jsYoz-Mn_ELgy6QvJSxpNcsdbEpM0EOCxFO2xq80bBZ2bmaPv67XmhKyANDMU0a92frdAD6PrVzfSI3vwcvtk"
                          />
                          <span className="font-bold">USD</span>
                        </div>
                      </div>
                    </div>
                    <div className="flex items-center justify-between text-sm text-text-light/70 dark:text-text-dark/70">
                      <div className="flex items-center gap-1">
                        <span className="material-symbols-outlined text-base text-secondary">
                          swap_vert
                        </span>
                        <span>1 USD = 0.92 EUR</span>
                      </div>
                      <div className="flex items-center gap-1">
                        <span className="material-symbols-outlined text-base text-secondary">
                          remove
                        </span>
                        <span>Fee: $2.50</span>
                      </div>
                    </div>
                    <div className="flex flex-col gap-2">
                      <label className="text-sm font-medium" htmlFor="receive-amount">
                        They get
                      </label>
                      <div className="flex items-center border border-border-light dark:border-border-dark rounded-lg overflow-hidden">
                        <input
                          className="flex-grow p-3 bg-neutral-light dark:bg-background-dark border-0"
                          id="receive-amount"
                          readOnly={true}
                          type="text"
                          value="917.50"
                        />
                        <div className="flex items-center gap-2 p-3 bg-neutral-light dark:bg-background-dark">
                          <img
                            alt="EU Flag"
                            className="h-5 rounded-sm"
                            src="https://lh3.googleusercontent.com/aida-public/AB6AXuBBtgY39skxiQ0f4uK451mQ_zIVB2mYP8oUT3HbdlqYRUHpXhHmYpzGcAOdBdncDfMii7OfDDHOB8iee0kF746ynMz4uH2MRRh9Wni6gDeTK3hhZIR-gEGd90XH_m5AiSiEhZEnBqKLYokXVqM6e19EhFDinVgz0LQXK8AHBgwWga5F_Wd6P-_00XeIkbfTbZOPbAQSVHXAZKHqw9x2Cl1ZKYG9dpNcJxzwfhbBbMnCYqqzwFuFWCs6ptZ7W779KTFqQslTDF4S_4I"
                          />
                          <span className="font-bold">EUR</span>
                        </div>
                      </div>
                    </div>
                    <button
                      onClick={() => navigate('/signup')}
                      className="w-full mt-2 flex cursor-pointer items-center justify-center overflow-hidden rounded-lg h-12 px-5 bg-secondary text-white text-base font-bold leading-normal tracking-[0.015em] hover:bg-secondary/90 transition-colors"
                    >
                      <span className="truncate">Get Started</span>
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>

        {/* FeatureSection */}
        <section className="w-full py-16 md:py-24 bg-neutral-light dark:bg-neutral-dark">
          <div className="max-w-7xl mx-auto px-6 md:px-10">
            <div className="flex flex-col gap-10">
              <div className="flex flex-col gap-4 text-center max-w-3xl mx-auto">
                <h2 className="text-3xl font-bold leading-tight tracking-[-0.033em] md:text-4xl">
                  Why Choose Globepay?
                </h2>
                <p className="text-base font-normal leading-normal text-text-light/80 dark:text-text-dark/80">
                  We offer a fast, secure, and affordable way to send money internationally,
                  connecting you with your loved ones without the hassle.
                </p>
              </div>
              <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                <div className="flex flex-col flex-1 gap-4 rounded-xl border border-border-light dark:border-border-dark bg-background-light dark:bg-background-dark p-6">
                  <div className="text-secondary">
                    <span className="material-symbols-outlined !text-4xl">rocket_launch</span>
                  </div>
                  <div className="flex flex-col gap-1">
                    <h3 className="text-lg font-bold leading-tight">Fast Transfers</h3>
                    <p className="text-sm font-normal leading-normal text-text-light/70 dark:text-text-dark/70">
                      Most transfers arrive within minutes, so your money gets where it needs to be,
                      faster.
                    </p>
                  </div>
                </div>
                <div className="flex flex-col flex-1 gap-4 rounded-xl border border-border-light dark:border-border-dark bg-background-light dark:bg-background-dark p-6">
                  <div className="text-secondary">
                    <span className="material-symbols-outlined !text-4xl">savings</span>
                  </div>
                  <div className="flex flex-col gap-1">
                    <h3 className="text-lg font-bold leading-tight">Low Fees</h3>
                    <p className="text-sm font-normal leading-normal text-text-light/70 dark:text-text-dark/70">
                      Enjoy transparent, competitive exchange rates and minimal fees on every
                      transaction.
                    </p>
                  </div>
                </div>
                <div className="flex flex-col flex-1 gap-4 rounded-xl border border-border-light dark:border-border-dark bg-background-light dark:bg-background-dark p-6">
                  <div className="text-secondary">
                    <span className="material-symbols-outlined !text-4xl">language</span>
                  </div>
                  <div className="flex flex-col gap-1">
                    <h3 className="text-lg font-bold leading-tight">Global Reach</h3>
                    <p className="text-sm font-normal leading-normal text-text-light/70 dark:text-text-dark/70">
                      Send money to over 150 countries with our ever-expanding network of partners.
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>

        {/* How It Works Section */}
        <section className="w-full py-16 md:py-24">
          <div className="max-w-7xl mx-auto px-6 md:px-10">
            <div className="flex flex-col gap-10">
              <div className="flex flex-col gap-4 text-center max-w-3xl mx-auto">
                <h2 className="text-3xl font-bold leading-tight tracking-[-0.033em] md:text-4xl">
                  How It Works
                </h2>
                <p className="text-base font-normal leading-normal text-text-light/80 dark:text-text-dark/80">
                  Get started in just a few simple steps.
                </p>
              </div>
              <div className="relative">
                <div
                  className="hidden md:block absolute top-1/2 left-0 w-full h-0.5 bg-border-light dark:bg-border-dark"
                  style={{ transform: 'translateY(-50%)' }}
                ></div>
                <div className="relative grid grid-cols-1 md:grid-cols-3 gap-8">
                  <div className="flex flex-col items-center text-center gap-4">
                    <div className="flex items-center justify-center size-16 rounded-full bg-primary/20 text-primary border-4 border-background-light dark:border-background-dark">
                      <span className="material-symbols-outlined !text-3xl">person_add</span>
                    </div>
                    <h3 className="text-lg font-bold">1. Create Your Account</h3>
                    <p className="text-sm text-text-light/70 dark:text-text-dark/70">
                      Sign up for free online or in our app. All you need is an email address.
                    </p>
                  </div>
                  <div className="flex flex-col items-center text-center gap-4">
                    <div className="flex items-center justify-center size-16 rounded-full bg-primary/20 text-primary border-4 border-background-light dark:border-background-dark">
                      <span className="material-symbols-outlined !text-3xl">add_circle</span>
                    </div>
                    <h3 className="text-lg font-bold">2. Enter Amount &amp; Recipient</h3>
                    <p className="text-sm text-text-light/70 dark:text-text-dark/70">
                      Tell us how much you want to send and your recipient&apos;s bank details.
                    </p>
                  </div>
                  <div className="flex flex-col items-center text-center gap-4">
                    <div className="flex items-center justify-center size-16 rounded-full bg-primary/20 text-primary border-4 border-background-light dark:border-background-dark">
                      <span className="material-symbols-outlined !text-3xl">send</span>
                    </div>
                    <h3 className="text-lg font-bold">3. Send &amp; Track</h3>
                    <p className="text-sm text-text-light/70 dark:text-text-dark/70">
                      Send your money and track its progress in real-time through our app.
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>

        {/* CTA Section */}
        <section className="w-full py-16 md:py-24 bg-primary text-white">
          <div className="max-w-7xl mx-auto px-6 md:px-10">
            <div className="flex flex-col gap-6 text-center max-w-3xl mx-auto">
              <h2 className="text-3xl font-bold leading-tight tracking-[-0.033em] md:text-4xl">
                Ready to Send Money Globally?
              </h2>
              <p className="text-base font-normal leading-normal text-white/90">
                Join thousands of satisfied customers who trust Globepay for their international
                money transfers.
              </p>
              <div className="flex flex-col sm:flex-row items-center justify-center gap-4 pt-4">
                <button
                  onClick={() => navigate('/signup')}
                  className="flex min-w-[84px] max-w-[240px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-12 px-5 bg-white text-primary text-base font-bold leading-normal tracking-[0.015em] hover:bg-white/90 transition-colors"
                >
                  <span className="truncate">Create Free Account</span>
                </button>
                <button
                  onClick={() => navigate('/login')}
                  className="flex min-w-[84px] max-w-[240px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-12 px-5 bg-transparent border border-white text-white text-base font-bold leading-normal tracking-[0.015em] hover:bg-white/10 transition-colors"
                >
                  <span className="truncate">Log In</span>
                </button>
              </div>
            </div>
          </div>
        </section>
      </main>

      {/* Footer */}
      <footer className="w-full py-10 bg-neutral-dark text-text-dark">
        <div className="max-w-7xl mx-auto px-6 md:px-10">
          <div className="flex flex-col md:flex-row items-center justify-between gap-6">
            <div className="flex items-center gap-4">
              <div className="size-8 text-white">
                <svg fill="none" viewBox="0 0 48 48" xmlns="http://www.w3.org/2000/svg">
                  <path
                    d="M42.4379 44C42.4379 44 36.0744 33.9038 41.1692 24C46.8624 12.9336 42.2078 4 42.2078 4L7.01134 4C7.01134 4 11.6577 12.932 5.96912 23.9969C0.876273 33.9029 7.27094 44 7.27094 44L42.4379 44Z"
                    fill="currentColor"
                  ></path>
                </svg>
              </div>
              <h2 className="text-xl font-bold leading-tight tracking-[-0.015em]">Globepay</h2>
            </div>
            <div className="flex flex-wrap items-center justify-center gap-6">
              <a
                className="text-sm font-medium leading-normal hover:text-secondary transition-colors"
                href="#"
              >
                Privacy Policy
              </a>
              <a
                className="text-sm font-medium leading-normal hover:text-secondary transition-colors"
                href="#"
              >
                Terms of Service
              </a>
              <a
                className="text-sm font-medium leading-normal hover:text-secondary transition-colors"
                href="#"
              >
                Security
              </a>
              <a
                className="text-sm font-medium leading-normal hover:text-secondary transition-colors"
                href="#"
              >
                Contact Us
              </a>
            </div>
            <p className="text-sm text-text-dark/70">© 2024 Globepay. All rights reserved.</p>
          </div>
        </div>
      </footer>
    </div>
  );
};

export default Welcome;
