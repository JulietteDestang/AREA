/*eslint-disable*/
import * as React from 'react'
import { Button, Dialog, DialogTitle, Box, TextField } from '@mui/material'
export default function TextInputsAParams ({ open, setOpen, newCard, setNewCard, fields }) {
    const [actionsFields, setActionFields] = React.useState()

    const handleSubmit = React.useCallback(() => {
        const temp = []
        actionsFields.forEach((action) => {
            temp.push(action)
        })
        setNewCard(newCard => ({ ...newCard, actionsFields: temp.join('@@@') }))
        setOpen(false)
    })

    React.useEffect(() => {
        let index = 0
        const temp = fields
        let args = { }
        const array = []
        for (let i = 0; i < temp.length; i++ ) {
            args = { id: index++, title: temp[i] }
            array.push(args)
            args = {}
        }
        setActionFields(array)
    }, [])

    return (
        <React.Fragment>
            <Dialog open={open}>
                <DialogTitle>Action parameters</DialogTitle>
                    <Box >
                        {actionsFields && actionsFields.map((field, index) => (
                            <React.Fragment>
                                <TextField
                                required
                                id={field.title}
                                label={field.title}
                                key={index}
                                onChange={(e) => setActionFields(actionsFields => {
                                    return [
                                        ...actionsFields.slice(0, index),
                                        actionsFields[index] = e.target.value,
                                        ...actionsFields.slice(index + 1)
                                    ]
                                })}
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
