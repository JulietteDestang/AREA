import * as React from 'react'
import { Button, Box, Dialog, DialogTitle, List, ListItemText, ListItem, FormControlLabel, FormGroup, Checkbox } from '@mui/material'
import { AREACard } from './Cards'
import './../App.css'
import TextInputsRParams from './TextInputsRParams'
import TextInputsAParams from './TextInputsAParams'
import NewAreaButton from './Icons/NewAreaButton'
import { createTheme, ThemeProvider, Typography } from '@material-ui/core'
import axios from 'axios'

export default function Wallet () {
    const [openDialog, setOpenDialog] = React.useState(false)
    const [singleCard, setSingleCard] = React.useState(false)
    const [areaCards, setAreaCards] = React.useState([])
    const [actionArray, setActionArray] = React.useState([])
    const [reactionArray, setReactionArray] = React.useState([])
    const [newCard, setNewCard] = React.useState({
        ID: null,
        action: null,
        actionsFields: null,
        actionService: null,
        reaction: null,
        reactionsFields: null,
        reactionService: null
    })
    const cards = []
    const actionData = []
    const reactionData = []

    React.useEffect(() => {
        axios.get('http://localhost:8080/area/user/areas', { withCredentials: true })
        .then(function (response) {
            const areas = response.data
            areas.forEach(area => {
                const formattedArea = {
                    ID: area.ID,
                    action: area.action_func,
                    actionService: area.action_service,
                    reaction: area.reaction_func,
                    reactionService: area.reaction_service
                }
                cards.unshift(formattedArea)
            })
            setAreaCards(cards)
        }).catch(function (error) {
            console.log(error)
        })
    }, [])

    React.useEffect(() => {
        axios.get('http://localhost:8080/area/user/propositions', { withCredentials: true })
        .then(function (response) {
            const services = response.data
            services.forEach(service => {
                const actionsArr = []
                const reactionsArr = []
                service.actions.forEach((item) => {
                    const fields = []
                    item.field_names.forEach((key) => {
                        fields.push(key)
                    })
                    const actionsObj = {
                        description: item.description,
                        fields
                    }
                    actionsArr.push(actionsObj)
                })
                const actionTotal = {
                    service: service.name,
                    actions: actionsArr
                }
                if (actionTotal.actions.length !== 0) {
                    actionData.push(actionTotal)
                }

                service.reactions.forEach((item) => {
                    const fields = []
                    item.field_names.forEach((key) => {
                        fields.push(key)
                    })
                    const reactionsObj = {
                        description: item.description,
                        fields
                    }
                    reactionsArr.push(reactionsObj)
                })
                const reactionTotal = {
                    service: service.name,
                    reactions: reactionsArr
                }
                if (reactionTotal.reactions.length !== 0) {
                    reactionData.push(reactionTotal)
                }
            })
            setActionArray(actionData)
            setReactionArray(reactionData)
        })
        .catch(function (error) {
            console.log(error)
        })
    }, [])

    const handleNewCard = () => {
        if (singleCard) {
            const headers = {
                'Content-Type': 'text/plain'
            }
            const sentCard = {
                action: newCard.action,
                actionService: newCard.actionService,
                actionFields: newCard.actionsFields,
                reaction: newCard.reaction,
                reactionService: newCard.reactionService,
                reactionFields: newCard.reactionsFields
            }
            axios.post('http://localhost:8080/area/create', {
                action_service: sentCard.actionService,
                action_func: sentCard.action,
                action_func_params: sentCard.actionFields,
                reaction_service: sentCard.reactionService,
                reaction_func: sentCard.reaction,
                reaction_func_params: sentCard.reactionFields
            }, { headers: { 'Content-Type': 'text/plain' }, withCredentials: true })
            .then(function (response) {
                sentCard.ID = response.data
                window.location.reload(true)
            })
            .catch(function (error) {
                console.log(error)
            })
            setNewCard({
                ID: null,
                action: null,
                actionService: null,
                reaction: null,
                reactionService: null
            })
            setSingleCard(false)
        }
        setOpenDialog(false)
    }

    const theme = createTheme({
        typography: {
          fontFamily: ['Titan One', 'cursive'].join(',')
        }
      })
    return (
        <React.Fragment>
                <Box sx={{
                    marginTop: 5,
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center'
                }}>
                    <ThemeProvider theme={theme}>
                        <Typography variant='h2' gutterBottom> AREA Wallet</Typography>
                    </ThemeProvider>
                    <Button size='small' onClick={ () => { setOpenDialog(true) } } className='newAreaButton'>
                        <NewAreaButton/>
                    </Button>
                </Box>
                <AREACard cards={areaCards} />
                <NewCardDialog
                    onClose={handleNewCard}
                    setSingleCard={setSingleCard}
                    singleCard={singleCard}
                    open={openDialog}
                    newCard={newCard}
                    setNewCard={setNewCard}
                    actionArray={actionArray}
                    reactionArray={reactionArray} />
        </React.Fragment >
    )
}

function NewCardDialog ({ setNewCard, newCard, actionArray, reactionArray, ...props }) {
    const [openAFieldsDialog, setOpenAFieldsDialog] = React.useState(false)
    const [openRFieldsDialog, setOpenRFieldsDialog] = React.useState(false)
    const [openServiceActionDialog, setOpenServiceActionDialog] = React.useState(false)
    const [openActionDialog, setOpenActionDialog] = React.useState(false)
    const [openServiceReactionDialog, setOpenServiceReactionDialog] = React.useState(false)
    const [openReactionDialog, setOpenReactionDialog] = React.useState(false)
    const [currentAction, setCurrentAction] = React.useState({
        description: null,
        fields: null
    })
    const [currentReaction, setCurrentReaction] = React.useState({
        description: null,
        fields: null
    })
    const [currentActionService, setCurrentActionService] = React.useState({
        service: null,
        actions: null,
        reactions: null
    })
    const [currentReactionService, setCurrentReactionService] = React.useState({
        service: null,
        actions: null,
        reactions: null
    })

    React.useEffect(() => {
        if (newCard.action != null && newCard.actionService != null && newCard.reaction != null && newCard.reactionService != null) {
            props.setSingleCard(true)
        }
    })

    const handleClickActionService = React.useCallback((service) => {
        setNewCard(newCard => ({ ...newCard, actionService: service.service }))
        setCurrentActionService(service)
        setOpenServiceActionDialog(false)
    }, [])
    const handleActionPick = React.useCallback((action) => {
        setNewCard(newCard => ({ ...newCard, action: action.description }))
        setCurrentAction(action)
        setOpenAFieldsDialog(true)
        setOpenActionDialog(false)
    }, [])

    const handleClickReactionService = React.useCallback((service) => {
        setNewCard(newCard => ({ ...newCard, reactionService: service.service }))
        setCurrentReactionService(service)
        setOpenServiceReactionDialog(false)
    }, [])
    const handleReactionPick = React.useCallback((reaction) => {
        setNewCard(newCard => ({ ...newCard, reaction: reaction.description }))
        setCurrentReaction(reaction)
        setOpenRFieldsDialog(true)
        setOpenReactionDialog(false)
    }, [])

    return (
        <React.Fragment>
            <Dialog onClose={props.onClose} open={props.open}>
                <DialogTitle>Create a new AREA :</DialogTitle>
                <FormGroup>
                    <FormControlLabel disabled control={<Checkbox checked={newCard.actionService !== null} />} label={<Button onClick={() => setOpenServiceActionDialog(true)}> {newCard.actionService ? newCard.actionService : 'Action Service'}</Button>} />
                    <FormControlLabel disabled control={<Checkbox checked={newCard.action !== null} />} label={<Button onClick={() => setOpenActionDialog(true)}> {newCard.action ? newCard.action : 'Action'}</Button>} />
                    <FormControlLabel disabled control={<Checkbox checked={newCard.reactionService !== null} />} label={<Button onClick={() => setOpenServiceReactionDialog(true)}> {newCard.reactionService ? newCard.reactionService : 'Reaction Service'}</Button>} />
                    <FormControlLabel disabled control={<Checkbox checked={newCard.reaction !== null} />} label={<Button onClick={() => setOpenReactionDialog(true)}> {newCard.reaction ? newCard.reaction : 'Reaction'}</Button>} />
                    <Button variant='outlined' disabled={!props.singleCard} onClick={() => { props.onClose(false) }}>Valider</Button>
                </FormGroup>
            </Dialog>

            <Dialog onClose={() => setOpenServiceActionDialog(false)} open={openServiceActionDialog}>
                <DialogTitle>Action Service</DialogTitle>
                    <List sx={{ pt: 0 }}>
                    {actionArray.map((service, index) => (
                        <ListItem button onClick={() => handleClickActionService(service) } key={index}>
                            <ListItemText primary={service.service} />
                    </ListItem>
                    ))}
                    </List>
            </Dialog>

            <Dialog onClose={() => setOpenActionDialog(false)} open={openActionDialog}>
                <DialogTitle>Action</DialogTitle>
                <List sx={{ pt: 0 }}>
                    { currentActionService?.actions && currentActionService.actions.map((action, index) => (
                        <ListItem button onClick={() => handleActionPick(action)} key={index}>
                            <ListItemText primary={action.description} />
                        </ListItem>
                    ))}
                </List>
            </Dialog >

            {currentAction?.fields &&
            <TextInputsAParams open={openAFieldsDialog} setOpen={setOpenAFieldsDialog} newCard={newCard} setNewCard={setNewCard} fields={currentAction?.fields}/>}

            <Dialog onClose={() => setOpenServiceReactionDialog(false)} open={openServiceReactionDialog}>
                <DialogTitle>Reaction Service</DialogTitle>
                    <List sx={{ pt: 0 }}>
                    {reactionArray.map((service, index) => (
                        <ListItem button onClick={() => handleClickReactionService(service) } key={index}>
                            <ListItemText primary={service.service} />
                    </ListItem>
                    ))}
                    </List>
            </Dialog>

            <Dialog onClose={() => setOpenReactionDialog(false)} open={openReactionDialog}>
                <DialogTitle>Reaction</DialogTitle>
                <List sx={{ pt: 0 }}>
                    { currentReactionService?.reactions && currentReactionService.reactions.map((reaction, index) => (
                        <ListItem button onClick={() => handleReactionPick(reaction)} key={index}>
                            <ListItemText primary={reaction.description} />
                        </ListItem>
                    ))}
                </List>
            </Dialog >

            {currentReaction?.fields &&
            <TextInputsRParams open={openRFieldsDialog} setOpen={setOpenRFieldsDialog} newCard={newCard} setNewCard={setNewCard} fields={currentReaction?.fields}/>}

        </React.Fragment >
    )
}
