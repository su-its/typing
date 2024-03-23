// components/buttons/LogoutButton.js
'use client';

import Link from 'next/link';
import { useRouter } from 'next/router';
import { Button } from "@chakra-ui/react";

const LogoutButton = () => {
    return (
        <Button as={ Link } href="login" colorScheme="blue" size="lg">
            Logout
        </Button>
    );
};

export default LogoutButton;
