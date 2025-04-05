import React from "react";
import { getCurrentUser } from "@/app/actions";
import type { User } from "@/types/user";
import styles from "@/assets/sass/molecules/UserCard.module.scss";

interface UserCardPresenterProps {
  user?: User;
}

export const UserCardPresenter = ({ user }: UserCardPresenterProps) => {
  return (
    <>
      <div className={styles["user-card"]}>
        <div className={styles.left}>
          <img src={"../../../img/user_default.png"} />
        </div>
        <div className={styles.right}>
          <div className={styles.name}>名前: {user ? user.handleName : "ログインしていません"}</div>
          <div className={styles.number}>学籍番号: {user ? user.studentNumber : "未ログイン"}</div>
        </div>
      </div>
    </>
  );
};

const UserCard = async (props?: UserCardPresenterProps) => {
  const user = await getCurrentUser();
  return <UserCardPresenter user={user} {...props} />;
};

export default UserCard;
