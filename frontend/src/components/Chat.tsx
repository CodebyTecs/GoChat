import React, { useEffect, useRef, useState } from 'react';
import type { Message } from '../types';

interface ChatProps {
    username: string;
}

const Chat: React.FC<ChatProps> = ({ username }) => {
    const [receiver, setReceiver] = useState('');
    const [text, setText] = useState('');
    const [messages, setMessages] = useState<Message[]>([]);
    const ws = useRef<WebSocket | null>(null);

    useEffect(() => {
        ws.current = new WebSocket('ws://localhost:8080/ws');

        ws.current.onmessage = (event) => {
            const msg: Message = JSON.parse(event.data);
            setMessages((prev) => [...prev, msg]);
        };

        ws.current.onclose = () => {
            console.log('WebSocket closed');
        };

        return () => {
            ws.current?.close();
        };
    }, []);

    const handleSend = () => {
        if (!text || !receiver) return;

        const msg: Message = {
            sender: username,
            receiver,
            text,
            created_at: new Date().toISOString(),
        };

        ws.current?.send(JSON.stringify(msg));
        setMessages((prev) => [...prev, msg]);
        setText('');
    };

    return (
        <div style={styles.container}>
            <h1 style={styles.title}>GoChat</h1>
            <input
                style={styles.input}
                placeholder="Получатель"
                value={receiver}
                onChange={(e) => setReceiver(e.target.value)}
            />
            <div style={styles.chat}>
                {messages.map((m, i) => (
                    <div key={i} style={m.sender === username ? styles.ownMessage : styles.message}>
                        <strong>{m.sender}:</strong> {m.text}
                    </div>
                ))}
            </div>
            <div style={styles.sendBox}>
                <input
                    style={styles.input}
                    placeholder="Сообщение..."
                    value={text}
                    onChange={(e) => setText(e.target.value)}
                />
                <button style={styles.button} onClick={handleSend}>Отправить</button>
            </div>
        </div>
    );
};

const styles: { [key: string]: React.CSSProperties } = {
    container: {
        padding: '2rem',
        fontFamily: 'Arial',
    },
    title: {
        color: '#00ADD8',
        fontSize: '2rem',
        marginBottom: '1rem',
    },
    chat: {
        border: '1px solid #ccc',
        padding: '1rem',
        height: '300px',
        overflowY: 'auto',
        marginBottom: '1rem',
        background: '#f9f9f9',
    },
    message: {
        padding: '0.5rem',
        marginBottom: '0.5rem',
        backgroundColor: '#e0e0e0',
        borderRadius: '5px',
    },
    ownMessage: {
        padding: '0.5rem',
        marginBottom: '0.5rem',
        backgroundColor: '#d0f0ff',
        borderRadius: '5px',
        alignSelf: 'flex-end',
    },
    input: {
        padding: '0.5rem',
        marginRight: '1rem',
        width: '200px',
    },
    sendBox: {
        display: 'flex',
        alignItems: 'center',
    },
    button: {
        padding: '0.5rem 1rem',
        backgroundColor: '#00ADD8',
        color: 'white',
        border: 'none',
        cursor: 'pointer',
    },
};

export default Chat;