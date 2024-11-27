import { MetaMaskInpageProvider } from '@metamask/sdk';

// Check if MetaMask is installed
const provider = window.ethereum; // Use window.ethereum directly

document.getElementById('signButton').onclick = async () => {
    if (!provider) {
        alert('Please install MetaMask!');
        return;
    }

    try {
        // Request account access
        const accounts = await provider.request({ method: 'eth_requestAccounts' });
        const address = accounts[0];

        // Get the challenge from user input
        const challenge = document.getElementById('challengeInput').value;

        if (!challenge) {
            alert('Please enter a challenge!');
            return;
        }

        // Sign the challenge
        const signature = await provider.request({
            method: 'personal_sign',
            params: [challenge, address],
        });

        // Output the result
        document.getElementById('output').innerText = `
            Address: ${address}
            Challenge: ${challenge}
            Signature: ${signature}
        `;
    } catch (error) {
        console.error(error);
        document.getElementById('output').innerText = 'Error signing message: ' + error.message;
    }
};
