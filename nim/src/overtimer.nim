import times, os, strutils, json, strformat

type
  EntryTypes {.pure} = enum
    overtime, spending

  Commands = enum
    start, stop, spend, status

  Entry = object
    entryType: EntryTypes
    startTime: DateTime

  FinishedEntry = object
    entryType: EntryTypes
    startTime: DateTime 
    endTime: DateTime 

  App = ref object
    currentEntry: ref Entry
    entries: seq[FinishedEntry]

proc execute(this: App, command: Commands) =
  echo command


proc getCommand(): Commands =
  let params = commandLineParams()
  if params.len() > 1:
    raise newException(ValueError, "only one parameter can be specified")
  elif params.len() == 0: 
    return Commands.status

  return parseEnum[Commands](params[0])

proc getApp(): App =
  let filename = "overtimer.json"
  if not fileExists(filename):
    echo fmt"{filename} not found, creating"
    writeFile(filename, "{}")

    return App()

  return readFile(filename).parseJson().to(App)


when isMainModule:
  let command = getCommand()
  let app = getApp()

  app.execute(command)

