
document.addEventListener('DOMContentLoaded', function() {
    const usernameInput = document.getElementById('username');
    const passwordInput = document.getElementById('password');
    const commandInput = document.getElementById('command');
    const loginBtn = document.getElementById('loginBtn');
    const pingBtn = document.getElementById('pingBtn');
    const sendBtn = document.getElementById('sendBtn');
    const outputDiv = document.getElementById('output');
    
    
    loginBtn.addEventListener('click', login);
    pingBtn.addEventListener('click', ping);
    sendBtn.addEventListener('click', sendCommand);
    
    
    commandInput.addEventListener('keypress', function(e) {
        if (e.key === 'Enter') {
            sendCommand();
        }
    });
    
    
    async function sendRequest(data) {
        try {
            outputDiv.textContent = 'Sending request...';
            
            const res = await fetch('/ami', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(data)
            });
            
            const result = await res.text();
            outputDiv.textContent = result;
        } catch (error) {
            outputDiv.textContent = `Error: ${error.message}`;
        }
    }
    
    // Action functions
    function login() {
        const username = usernameInput.value;
        const password = passwordInput.value;
        
        if (!username || !password) {
            outputDiv.textContent = 'Error: Username and password are required';
            return;
        }
        
        sendRequest({ action: 'login', username, secret: password });
    }
    
    function ping() {
        const username = usernameInput.value;
        const password = passwordInput.value;
        
        if (!username || !password) {
            outputDiv.textContent = 'Error: Username and password are required';
            return;
        }
        
        sendRequest({ action: 'ping', username, secret: password });
    }
    
    function sendCommand() {
        const username = usernameInput.value;
        const password = passwordInput.value;
        const command = commandInput.value;
        
        if (!username || !password) {
            outputDiv.textContent = 'Error: Username and password are required';
            return;
        }
        
        if (!command) {
            outputDiv.textContent = 'Error: Command is required';
            return;
        }
        
        sendRequest({ action: 'command', command, username, secret: password });
    }
});