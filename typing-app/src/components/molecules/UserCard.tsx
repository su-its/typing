import React from "react";
import type { StackProps } from "@chakra-ui/react";
import { getCurrentUser } from "@/app/actions";
import type { User } from "@/types/user";
import styles from "@/assets/sass/molecules/UserCard.module.scss";

interface UserCardPresenterProps extends StackProps {
  user?: User;
}

export const UserCardPresenter = ({ user, ...rest }: UserCardPresenterProps) => {
  const props: StackProps = {
    width: rest?.width ?? "18%",
    ...rest,
  };

  return (
    <>
      <div className={styles["user-card"]}>
        <div className={styles.left}>
          <img src={"https://www.shizuoka.ac.jp/cms/files/shizudai/MASTER/0100/uISrbYCb_VL033_r03.png"} />
        </div>
        <div className={styles.right}>
          <div className={styles.name}>{user ? user.handleName : "ログインしていません"}</div>
          <div className={styles.number}>学籍番号: {user ? user.studentNumber : ""}</div>
        </div>
      </div>
    </>
  );
};

const UserCard = async (props?: StackProps) => {
  const user = await getCurrentUser();
  return <UserCardPresenter user={user} {...props} />;
};

export default UserCard;
