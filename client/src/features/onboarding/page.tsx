import { useContext, useEffect } from "react";
import { styles } from "./page.css";
import { LoadingContext } from "@/providers/loading";

export const OnboardingPage = () => {
  const { setIsLoading } = useContext(LoadingContext);

  useEffect(() => {
    setIsLoading(true);

    setTimeout(() => {
      setIsLoading(false);
    }, 3000);
  }, [setIsLoading]);

  return (
    <div className={styles.container}>
      <img src="/maximum.svg" alt="Maximum" width="300" height="100" />
      <h1 className={styles.title}>MELINE</h1>
      <p className={styles.subtitle}>Coming Soon...</p>
    </div>
  );
};
