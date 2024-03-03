import React, {useEffect, useState} from "react";
import UserTimer from "./UserTimer";
import './App.css';

//export const serverBaseUrl = "http://localhost:8080";
export const serverBaseUrl = "";

const App = () => {
    const [data, setData] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch(`${serverBaseUrl}/v1/user`);

                if (response.status === 200) {
                    const parsedData = await response.json();
                    setData(parsedData)
                    //console.log(parsedData)
                }
            } catch (error) {
                console.log("Error fetching data: ", error);
            }
        };

        const intervalId = setInterval(fetchData, 1000);
        return () => clearInterval(intervalId);
    }, []);

    const handleAddUser = async (e, userName) => {
        if (e.code === "Enter") {
            await fetch(`${serverBaseUrl}/v1/user`,
                {method: "POST", body: JSON.stringify({name: userName})})
                .then(response => response.json())
                .then(data => console.log('Response:', data))
                .catch(error => console.error('Error:', error));
        }
    }

    return (
        <div className="App">
            <div className="App-header">
                <h2>Stopwatch</h2>
            </div>
            <div className="App-intro">
                {data.map(item => (
                    <UserTimer key={item.id} id={item.id} name={item.name} speaking={item.speaking}
                               speakDuration={item.speakDuration}/>
                ))}

                <div className="watch">
                    <div className="userAndControls">
                        <div className="username">
                            New User
                        </div>
                    </div>
                    <div className="numbers">
                        <div className="input-underlined">
                            <input id="userLabel" required onKeyDown={(e) => handleAddUser(e, e.target.value)}/>
                        </div>
                    </div>
                    <label htmlFor="userLabel">Hit Enter to submit</label>
                </div>
            </div>
        </div>
    );
}

export default App;
