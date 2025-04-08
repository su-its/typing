import React from "react";
import BrandText from "../molecules/BrandText";
import Separator from "../atoms/Separater";
import styles from "@/assets/sass/organism/Footer.module.scss";

const Footer: React.FC = () => {
  return (
    <>
      <div className={styles.footer}>
        <Separator />
        <div className={styles["brand-text"]}>
          <BrandText />
        </div>
      </div>
    </>
  );
};

export default Footer;
