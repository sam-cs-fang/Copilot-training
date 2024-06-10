

export const signup = async (username: string, password: string): Promise<void> => {
    const response = await fetch('http://localhost:6741/api/v1/users/signup', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password }),
    });

    if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
    }

    const token = await response.json();
    localStorage.setItem('ai-training-token', token);
}

export const signin = async (username: string, password: string): Promise<void> => {
    const response = await fetch('http://localhost:6741/api/v1/users/signin', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password }),
    });

    if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
    }

    const token = await response.json();
    console.log(token);
    localStorage.setItem('ai-training-token', token.data);
}