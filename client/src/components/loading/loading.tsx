import { styles } from './loading.css';
import { LoadingContext } from '@/providers/loading';
import React, { useContext } from 'react';

export const Loading= () => {
  const { isLoading } = useContext(LoadingContext);

  if (!isLoading) {
    return null;
  }

  return (
    <div className={styles.overlay}>
      <div className={styles.loading} />
    </div>
  );
};
