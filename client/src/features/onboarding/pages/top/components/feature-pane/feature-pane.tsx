import { MaximumIcon } from "@/components/icons/maximum";
import { styles } from "./feature-pane.css";

export const FeaturePane = () => {
  return (
    <div className={styles.container}>
      <div className={styles.icon}>
        <MaximumIcon />
      </div>
      <div className={styles.content}>
        <h1 className={styles.title}>MELINE</h1>
        <p className={styles.description}>
          埼玉大学プログラミングサークル Maximum の<br />
          コミュニケーションプラットフォーム
        </p>
      </div>
    </div>
  );
};
