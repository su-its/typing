import React from "react";
import styles from "@/assets/sass/molecules/BrandText.module.scss";
import brandImage from "@/assets/images/brand.png";

const BrandText: React.FC = () => {
  return (
    <div className={styles["brand-text"]}>
      {/* eslint-disable-next-line @next/next/no-img-element*/}
      <img src={brandImage.src} alt="Brand" />
    </div>
  );
};

export default BrandText;
