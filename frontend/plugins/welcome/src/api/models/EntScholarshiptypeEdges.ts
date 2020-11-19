/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntScholarshipRequest,
    EntScholarshipRequestFromJSON,
    EntScholarshipRequestFromJSONTyped,
    EntScholarshipRequestToJSON,
    EntScholarshipinformation,
    EntScholarshipinformationFromJSON,
    EntScholarshipinformationFromJSONTyped,
    EntScholarshipinformationToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntScholarshiptypeEdges
 */
export interface EntScholarshiptypeEdges {
    /**
     * ScholarshipRequest holds the value of the ScholarshipRequest edge.
     * @type {Array<EntScholarshipRequest>}
     * @memberof EntScholarshiptypeEdges
     */
    scholarshipRequest?: Array<EntScholarshipRequest>;
    /**
     * Scholarshipinformation holds the value of the Scholarshipinformation edge.
     * @type {Array<EntScholarshipinformation>}
     * @memberof EntScholarshiptypeEdges
     */
    scholarshipinformation?: Array<EntScholarshipinformation>;
}

export function EntScholarshiptypeEdgesFromJSON(json: any): EntScholarshiptypeEdges {
    return EntScholarshiptypeEdgesFromJSONTyped(json, false);
}

export function EntScholarshiptypeEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntScholarshiptypeEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'scholarshipRequest': !exists(json, 'scholarshipRequest') ? undefined : ((json['scholarshipRequest'] as Array<any>).map(EntScholarshipRequestFromJSON)),
        'scholarshipinformation': !exists(json, 'scholarshipinformation') ? undefined : ((json['scholarshipinformation'] as Array<any>).map(EntScholarshipinformationFromJSON)),
    };
}

export function EntScholarshiptypeEdgesToJSON(value?: EntScholarshiptypeEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'scholarshipRequest': value.scholarshipRequest === undefined ? undefined : ((value.scholarshipRequest as Array<any>).map(EntScholarshipRequestToJSON)),
        'scholarshipinformation': value.scholarshipinformation === undefined ? undefined : ((value.scholarshipinformation as Array<any>).map(EntScholarshipinformationToJSON)),
    };
}


