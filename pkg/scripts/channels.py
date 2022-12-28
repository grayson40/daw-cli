# Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
import json

# Creates sampler dictionary
Sampler = lambda name, channel: {
    "name": name,
    "channel": channel,
}

# Creates instrument dictionary
Instrument = lambda name, channel, plugin: {
    "name": name,
    "channel": channel,
    "plugin": plugin,
}

# Creates channel rack dictionary
ChannelRack = lambda samplers, instruments: {
    "samplers": samplers,
    "instruments": instruments,
}

# Writes project channels to json
def getChannels(project):
    # Get project channel rack
    project_channel_rack = project.channels

    # Get samplers
    samplers = project_channel_rack.samplers

    # Populate samplers list
    samplers_list = []
    samplers_iter = iter(samplers)
    for sample in samplers_iter:
        samplers_list.append(Sampler(name=sample.name, channel=sample.iid))

    # Get instruments
    instruments = project_channel_rack.instruments

    # Populate instruments list
    instruments_list = []
    instruments_iter = iter(instruments)
    for instrument in instruments_iter:
        instruments_list.append(
            Instrument(name=instrument.name, channel=instrument.iid, plugin=instrument.plugin.name)
        )

    # Create channel rack dictionary
    channel_rack_dict = ChannelRack(samplers=samplers_list, instruments=instruments_list)

    # Serializing json
    channel_rack_json = json.dumps(channel_rack_dict, indent=4)

    # Write to json
    with open("./project/channel-rack.json", "w") as outfile:
        outfile.write(channel_rack_json)
