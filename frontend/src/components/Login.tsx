import React, { useState } from 'react';
import { registerUser } from '../services/api';
import type { User } from '../services/api';

interface LoginProps {
    onLogin: (username: string) => void;
}

const Login: React.FC<LoginProps> = ({ onLogin }) => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState<string | null>(null);

    const handleRegister = async () => {
        try {
            const user: User = { username, password };
            await registerUser(user);
            onLogin(username); // ✅ используем коллбэ
        } catch (err) {
            setError('Ошибка регистрации');
            console.error(err);
        }
    };

    const handleLogin = () => {
        if (username && password) {
            onLogin(username); // ✅ используем коллбэк
        } else {
            setError('Введите логин и пароль');
        }
    };

    return (
        <div style={styles.container}>
            <h1 style={styles.title}>GoChat</h1>
            <div style={styles.form}>
                <input
                    style={styles.input}
                    type="text"
                    placeholder="Логин"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                />
                <input
                    style={styles.input}
                    type="password"
                    placeholder="Пароль"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
                <div style={styles.buttons}>
                    <button style={styles.button} onClick={handleLogin}>Войти</button>
                    <button style={styles.button} onClick={handleRegister}>Регистрация</button>
                </div>
                {error && <p style={{ color: 'red' }}>{error}</p>}
            </div>
        </div>
    );
};

const styles: { [key: string]: React.CSSProperties } = {
    container: {
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        justifyContent: 'center',
        height: '100vh',
        backgroundColor: '#ffffff',
    },
    title: {
        color: '#00ADD8',
        fontSize: '3rem',
        marginBottom: '2rem',
    },
    form: {
        display: 'flex',
        flexDirection: 'column',
        width: '300px',
    },
    input: {
        marginBottom: '1rem',
        padding: '0.5rem',
        fontSize: '1rem',
    },
    buttons: {
        display: 'flex',
        justifyContent: 'space-between',
    },
    button: {
        padding: '0.5rem 1rem',
        fontSize: '1rem',
        backgroundColor: '#00ADD8',
        color: 'white',
        border: 'none',
        cursor: 'pointer',
    },
};

export default Login;