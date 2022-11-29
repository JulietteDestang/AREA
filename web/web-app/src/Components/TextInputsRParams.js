/*eslint-disable*/
import * as React from 'react'
import { Button, Box, Dialog, DialogTitle, TextField } from '@mui/material'

export default function TextInputsRParams ({ open, setOpen, newCard, setNewCard, fields }) {
    const [reactionsFields, setReactionFields] = React.useState()

    const handleSubmit = React.useCallback(() => {
        const temp = []
        reactionsFields.forEach((reaction) => {
            temp.push(reaction)
        })
        setNewCard(newCard => ({ ...newCard, reactionsFields: temp.join('@@@') }))
        setOpen(false)
    })

    React.useEffect(() => {
        let index = 0
        const temp = fields
        let args = { }
        const array = []
        for (let i = 0; i < temp.length; i++) {
            args = { id: index++, title: temp[i] }
            array.push(args)
            args = {}
        }
        setReactionFields(array)
    }, [])

    return (
        <React.Fragment>
            <Dialog open={open}>
                <DialogTitle>Reaction parameters</DialogTitle>
                    <Box >
                        {reactionsFields && reactionsFields.map((field, index) => (
                            <React.Fragment>
                                <TextField
                                required
                                id={field.title}
                                label={field.title}
                                key={index}
                                onChange={ (e) => setReactionFields(reactionsFields => {
                                    return [
                                        ...reactionsFields.slice(0, index),
                                        reactionsFields[index] = e.target.value,
                                        ...reactionsFields.slice(index + 1)
                                    ]}) }
                                />
                            </React.Fragment>
                        ))}
                        <Button
                        type='submit'
                        fullwidth='true'
                        variant='contained'
                        onClick={handleSubmit}
                        >Submit</Button>
                </Box>
            </Dialog>
        </React.Fragment>
    )
}
