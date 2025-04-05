import React from "react";
// import { useAuth } from "@/hooks/useAuth";　// TODO: 実装
import Banner from "@/components/atoms/Banner";
import UserCard from "@/components/molecules/UserCard";
import Separator from "@/components/atoms/Separater";
import styles from "@/assets/sass/organism/Header.module.scss";

const Header: React.FC = () => {
  return (
    <>
      <div className={styles.header}>
        <Banner />
        <UserCard />
      </div>
      <Separator />
    </>
  );
};

export default Header;
