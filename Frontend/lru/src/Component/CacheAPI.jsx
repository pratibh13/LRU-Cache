import React, { useState } from 'react';
import axios from 'axios';
 
const CacheAPI = () => {
    const [keyInput, setKeyInput] = useState('');
    const [valueInput, setValueInput] = useState('');
    const [response, setResponse] = useState('');
    const [expirationTimeInput, setExpirationTimeInput] = useState('');
 
    const handleGetKey = async () => {
        try {
const res = await axios.get(`http://localhost:8080/get?key=${keyInput}`);
            if (res.data) {
                setResponse(`Key: ${keyInput}, Value: ${res.data}`);
            }
        } catch (error) {
            console.error(error);
            setResponse(`Key: ${keyInput} has been removed from the cache Or key does not exist`);
        }
    };
 
    const handleSetKey = async () => {
        try {
            const expirationSeconds=parseInt(expirationTimeInput)
            const expiration=expirationSeconds*1000
            console.log({ key: keyInput, value: valueInput,expirationTime: expiration })
await axios.post(`http://localhost:8080/set`, { key: keyInput, value: valueInput,expirationTime: expiration });
            setResponse(`Key ${keyInput} set successfully!`);
        } catch (error) {
            console.error(error);
            setResponse('Error setting key');
        }
    };
 
    return (
        <div>
            <div>
                <label>Key:</label>
                <input type="text" value={keyInput} onChange={(e) => setKeyInput(e.target.value)} />
                <button onClick={handleGetKey}>Get Key</button>
            </div>
            <div>
                <label>Value:</label>
                <input type="text" value={valueInput} onChange={(e) => setValueInput(e.target.value)} />
                <button onClick={handleSetKey}>Set Key</button>
            </div>
            <div>
                <label>Expiration:</label>
                <input type="number" value={expirationTimeInput} onChange={(e) => setExpirationTimeInput(e.target.value)} />
                
            </div>
            {response && <p>{response}</p>}
        </div>
    );
};
 
export default CacheAPI;