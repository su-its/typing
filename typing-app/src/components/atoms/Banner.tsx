import React from "react";
import Link from "next/link";
import styles from "@/assets/sass/atoms/Banner.module.scss";
import bannerImage from "@/assets/images/banner.png";

const Banner: React.FC = () => {
  return (
    <div className={styles.banner}>
      <Link href="/">
        {/* eslint-disable-next-line @next/next/no-img-element */}
        <img src={bannerImage.src} alt="Logo" />
      </Link>
    </div>
  );
};

export default Banner;
