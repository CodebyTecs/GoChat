import { useState } from 'react';
import Login from './components/Login';
import Chat from './components/Chat';

function App() {
    const [loggedInUser, setLoggedInUser] = useState<string | null>(null);

    return (
        <div>
            {loggedInUser ? (
                <Chat username={loggedInUser} />
            ) : (
                <Login onLogin={(username) => setLoggedInUser(username)} />
            )}
        </div>
    );
}

export default App;
