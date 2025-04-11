import React from "react";
import BrandText from "../molecules/BrandText";
import Separator from "../atoms/Separater";
import styles from "@/assets/sass/organism/Footer.module.scss";

const date = new Date();

const Footer: React.FC = () => {
  return (
    <div className={styles.footer}>
      <Separator />
      <BrandText />
      <div className={styles.text}>
        (c) 2024-{date.getFullYear()} Faculty of Informatics, Shizuoka University all rights reserved. Developed by IT
        Solution Room, Shizuoka University.
      </div>
    </div>
  );
};

export default Footer;
