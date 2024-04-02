import React from 'react';
import { IconButton } from '@chakra-ui/react';

const RefreshIcon = () => (
    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <polyline points="23 4 23 10 17 10"></polyline>
        <polyline points="1 20 1 14 7 14"></polyline>
        <path d="M3.51 9a9 9 0 0 1 14.85-3.36L23 10"></path>
        <path d="M20.49 15a9 9 0 0 1-14.85 3.36L1 14"></path>
    </svg>

);

interface RefreshButtonProps {
    onClick: () => void;
    isDisabled?: boolean;
}

const RefreshButton: React.FC<RefreshButtonProps> = ({ onClick, isDisabled = false }) => (
    <IconButton
        width={"50px"}
        right={"250px"}
        bg="#2B6CB0"
        color="White"
        aria-label="refresh"
        icon={<RefreshIcon />}
        onClick={onClick}
        isDisabled={isDisabled}
    />
);

export default RefreshButton;


