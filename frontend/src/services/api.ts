const API_URL = 'http://localhost:50051'; // если grpc-gateway, иначе укажи REST proxy

export interface User {
    username: string;
    password: string;
}

export interface Message {
    sender: string;
    receiver: string;
    text: string;
    created_at?: string;
}

// Регистрация пользователя
export async function registerUser(user: User): Promise<void> {
    const res = await fetch(`${API_URL}/RegisterUser`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(user),
    });

    if (!res.ok) {
        throw new Error('Registration failed');
    }
}

// Отправка сообщения
export async function sendMessage(message: Message): Promise<void> {
    const res = await fetch(`${API_URL}/SendMessage`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(message),
    });

    if (!res.ok) {
        throw new Error('Message sending failed');
    }
}