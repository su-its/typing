// components/buttons/LogoutButton.js
'use client';

import { useRouter } from 'next/router';
import { Button } from "@chakra-ui/react";

const LogoutButton = () => {
    return (
        <Button colorScheme="blue" size="lg">
            Logout
        </Button>
    );
};

export default LogoutButton;
