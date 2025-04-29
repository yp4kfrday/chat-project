import {useState} from 'react';
import './App.css'

import Chat from "./components/Chat";

function App() {
	const [username, setUsername] = useState("");
	const [isConnected, setIsConnected] = useState(false);

	const handleConnect = () => {
		if (username.trim()) {
			setIsConnected(true);
		}
	};

	return (
		<div className="app">
			{!isConnected ? (
				<div className="login">
					<h2>Введите имя</h2>
					<input
						type="text"
						value={username}
						onChange={(e) => setUsername(e.target.value)}
						placeholder="Ваше имя..."
					/>
					<button onClick={handleConnect}>Войти</button>
				</div>
			) : (
				<Chat username={username} />
			)}
		</div>
	);
}


export default App
