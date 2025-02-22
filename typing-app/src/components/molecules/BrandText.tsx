import React from "react";
import styles from "@/assets/sass/molecules/BrandText.module.scss";
import brandImage from "@/assets/images/brand.png";

const BrandText: React.FC = () => {
  return (
    <div className={styles["brand-text"]}>
      <img src={brandImage.src} />
    </div>
  );
};

export default BrandText;
