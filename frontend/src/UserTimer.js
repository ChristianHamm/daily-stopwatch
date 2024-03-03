import React from 'react'
import './UserTimer.css';

//export const serverBaseUrl = "http://localhost:8080";
export const serverBaseUrl = "";

const UserTimer = props => {
    const formattedTimer = formatTimer(props.speakDuration)

    const requestOptions = {
        method: 'PUT',
        headers: {'Content-Type': 'application/json',},
    };

    const toggleSpeaking = async (id) => {
        console.log("toggled speaking: ")
        await fetch(`${serverBaseUrl}/v1/user/` + id, requestOptions)
            .then(response => response.json())
            .then(data => console.log('Response:', data))
            .catch(error => console.error('Error:', error));
    }

    const toggleReset = async (id) => {
        console.log("toggled reset: ")
        await fetch(`${serverBaseUrl}/v1/user/` + id + `?reset=true`, requestOptions)
            .then(response => response.json())
            .then(data => console.log('Response:', data))
            .catch(error => console.error('Error:', error));
    }

    const deleteUser = async (id) => {
        console.log("delete user: ")
        await fetch(`${serverBaseUrl}/v1/user/` + id,
            {method: "DELETE"})
            .then(response => response.json())
            .then(data => console.log('Response:', data))
            .catch(error => console.error('Error:', error));
    }

    return (
        <div className="watch">
            <div className="userAndControls">
                <div className="resetButton" onClick={() => toggleReset(props.id)}>
                    <i className="fa fa-refresh"/>
                </div>
                <div className="username">{props.name}</div>
                <div className="deleteButton" onClick={() => deleteUser(props.id)}>
                    <i className="fa fa-remove"/>
                </div>
            </div>
            <div className={props.speaking ? "numbers numbersStarted" : "numbers"}
                 onClick={() => toggleSpeaking(props.id)}>{formattedTimer}</div>
        </div>
    )
}

function formatTimer(timer) {
    if (isNaN(parseFloat(timer))) {
        return "00:00"
    }

    let timeRepr, secondsRepr, minutesRepr, hoursRepr;
    let seconds = timer % 60;
    let minutes = Math.floor(timer / 60) % 60;
    let hours = Math.floor(timer / 3600) % 24;

    if (seconds < 10)
        secondsRepr = "0" + seconds.toString();
    else secondsRepr = seconds;

    if (minutes < 10) minutesRepr = "0" + minutes.toString();
    else minutesRepr = minutes;

    if (hours < 10) hoursRepr = "0" + hours.toString();
    else hoursRepr = hours;

    if (hours > 0)
        timeRepr = hoursRepr + ":" + minutesRepr + ":" + secondsRepr;
    else
        timeRepr = minutesRepr + ":" + secondsRepr;

    return timeRepr;
}

UserTimer.propTypes = {}
export default UserTimer
