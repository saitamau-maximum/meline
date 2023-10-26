import { lazy } from "react";
import { Route, Routes } from "react-router-dom";

const Top = lazy(() => import("./pages/top"));

export const OnBoardingRoutes = () => {
  return (
    <Routes>
      <Route path="/" element={<Top />} />
    </Routes>
  );
};
