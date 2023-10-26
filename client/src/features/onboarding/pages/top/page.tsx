import { styles } from "./page.css";

export const OnboardingPage = () => {
  return (
    <div className={styles.container}>
      <img src="/maximum.svg" alt="Maximum" width="300" height="100" />
      <h1 className={styles.title}>MELINE</h1>
      <p className={styles.subtitle}>Coming Soon...</p>
    </div>
  );
};
