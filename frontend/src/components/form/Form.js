import React from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import { Button } from 'react-bootstrap';
import { Form, Container, Row, Col } from 'react-bootstrap';


class FormSubmission extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            to: '',
            from_name: '',
            subject: '',
            content: '',
            mail_status: '',
        }
        this.handleChange = this.handleChange.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)
    }

    handleChange(event) {
        switch (event.target.name) {
            case "to":
                this.setState({ to: event.target.value })
                return
            case "from_name":
                this.setState({ from_name: event.target.value })
                return
            case "subject":
                this.setState({ subject: event.target.value })
                return
            case "content":
                this.setState({ content: event.target.value })
                return
            default:
                return
        }
    }

    handleSubmit(event) {
        event.preventDefault();
        const mail = {to: this.state.to, from_name: this.state.from_name, subject: this.state.subject, content: this.state.content}
        const requestOptions = {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(mail)
        };
    
        fetch('/mail/send', requestOptions)
            .then((response) => {
                return response.json()

            })
            .then((data) => {
                this.setState({to: '', from_name: '', subject: '', content: '', mail_status: 'Email sent'})
                console.log(data)
            })
            .catch(err => console.log(err))
        }
    
    buttonStyle = () => {
        return {minWidth: "300px"}
    }

    render() {
        return(
            <Container style={{marginTop: "20px"}}>
                <form onSubmit={this.handleSubmit}>
                        <Row>
                            <Form.Group className="mb-3">
                                <Form.Control name="to" value={this.state.to} onChange={this.handleChange} type="email" placeholder="Enter email" />
                            </Form.Group>
                        </Row>
                        <Row>
                            <Col>
                                <Form.Group className="mb-3">
                                    <Form.Control name="from_name" value={this.state.from_name} onChange={this.handleChange} type="text" placeholder="Senders name" />
                                </Form.Group>
                            </Col>
                            <Col>
                                <Form.Group className="mb-3">
                                    <Form.Control name="subject" value={this.state.subject} onChange={this.handleChange} type="text" placeholder="Subject" />
                                </Form.Group>
                            </Col>  
                        </Row>
                        <Row>
                            <Col>
                                <Form.Group className="mb-3">
                                    <Form.Control name="content" value={this.state.content} onChange={this.handleChange} type="text" placeholder="Content" />
                                </Form.Group>
                            </Col>
                        </Row>
                        <Row>
                            <Col>
                                <Button type="submit" style={this.buttonStyle()}>Send</Button>
                            </Col>
                        </Row>
                </form>
                <div>{this.state.mail_status}</div>
            </Container>
        )
    }
}
export default FormSubmission;