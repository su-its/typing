// components/buttons/LogoutButton.js
'use client';

import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { Button } from "@chakra-ui/react";

const LogoutButton = () => {
    const router = useRouter();

    const handleLogout = async () => {
        //TODO:ログアウト処理を実装
        router.push('/login');
    }

    return (
        <Button colorScheme="blue" size="lg" onClick={handleLogout}>
            Logout
        </Button>
    );
};

export default LogoutButton;
