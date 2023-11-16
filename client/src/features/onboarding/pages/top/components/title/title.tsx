import { styles } from "./title.css";

export const Title = () => {
  return (
    <div className={styles.container}>
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
