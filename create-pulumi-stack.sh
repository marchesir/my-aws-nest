#!/usr/bin/env bash
# create-pulumi-stack.sh
# Generates a Docker-style stack name and creates a Pulumi stack in a folder with that name.

# TODO: Pass AWS type and generate minimal Pulumi stack config for it.

# --- Docker-style name generator ---
ADJECTIVES=(happy sleepy brave clever tiny fast quiet)
ANIMALS=(otter fox tiger panda eagle shark rabbit)
RAND_ADJ=${ADJECTIVES[$RANDOM % ${#ADJECTIVES[@]}]}
RAND_ANIMAL=${ANIMALS[$RANDOM % ${#ANIMALS[@]}]}
STACK_NAME="${RAND_ADJ}${RAND_ANIMAL}"

echo "Creating Pulumi stack: $STACK_NAME"

# --- Create Pulumi stack ---
pulumi stack init "$STACK_NAME"

echo "Stack $STACK_NAME created: the stack is only initialized locally: to persist commit and run 'pulumi up' to deploy."